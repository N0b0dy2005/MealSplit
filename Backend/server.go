package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"github.com/vektah/gqlparser/v2/ast"
	"log"
	"my-backend/controller"
	"my-backend/graph"
	"my-backend/utils/cast"
	"net/http"
	"os"
	"time"
)

const (
	defaultPort = "8080"
	secretKey   = "mysecret" // In der Produktion sicher aufbewahren!
)

// User-Struktur aus der Datenbank
type User struct {
	ID       int            `db:"id"`
	Email    sql.NullString `db:"email"`
	Password sql.NullString `db:"password"`
}

// loginRequest repräsentiert die JSON-Daten, die beim Login erwartet werden.
type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Aufbau der Datenbankverbindung mit sqlx
	dsn := "root:felix123@tcp(127.0.0.1:3306)/MealSplit"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("Datenbankverbindungsfehler:", err)
	}
	defer db.Close()

	// Übergabe der DB-Verbindung an den Resolver
	resolver := &graph.Resolver{}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// Einrichtung des HTTP-Mux mit geschützten und ungeschützten Routen
	mux := http.NewServeMux()
	// Unprotected Login-Route
	mux.HandleFunc("/api/login", loginHandler(db))
	// Alle weiteren Routen werden durch die Auth-Middleware geschützt
	mux.Handle("/api/playground", authMiddleware(playground.Handler("GraphQL playground", "/query"), db))
	mux.Handle("/api/query", authMiddleware(srv, db))

	log.Printf("Server läuft auf http://localhost:%s/", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalln(err)
	}
}

// loginHandler verarbeitet POST-Anfragen zur Anmeldung.
func loginHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Nur POST-Anfragen zulassen
		if r.Method != http.MethodPost {
			http.Error(w, "Methode nicht erlaubt", http.StatusMethodNotAllowed)
			return
		}

		// ParseForm aufrufen, um Form-Daten zu lesen
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Ungültige Anfrage", http.StatusBadRequest)
			return
		}

		// Form-Werte auslesen

		req := loginRequest{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		// Abrufen des Benutzers anhand der E-Mail
		var user User
		err := db.Get(&user, "SELECT id, email, password  FROM users WHERE email = ?", req.Email)
		if err != nil {
			http.Error(w, "Ungültige Anmeldedaten", http.StatusUnauthorized)
			return
		}

		// hash req.Password as  sha256.New in 64 bytes
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password)))

		// In your loginHandler function, add these lines:
		fmt.Println("Provided email:", req.Email)
		fmt.Println("Generated hash:", hash)
		fmt.Println("Stored password:", user.Password.String)
		fmt.Println("Password valid?", user.Password.Valid)
		// Überprüfung des Passworts (Hinweis: In einer echten Anwendung sollte das Passwort gehasht sein)
		if !user.Password.Valid || user.Password.String != hash {
			http.Error(w, "Ungültige Anmeldedaten", http.StatusUnauthorized)
			return
		}

		// Erstellen eines JWT-Tokens mit einer Gültigkeit von 24 Stunden
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"exp":     time.Now().Add(24 * time.Hour).Unix(),
		})
		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			http.Error(w, "Fehler beim Erstellen des Tokens", http.StatusInternalServerError)
			return
		}

		// Setzen des JWT als HTTP‑Only Cookie
		cookie := &http.Cookie{
			Name:     "jwt",
			Value:    tokenString,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)

		// Rückmeldung an den Client
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login erfolgreich"))
	}
}

// authMiddleware prüft bei jeder Anfrage, ob ein gültiges JWT im Cookie vorhanden ist.
func authMiddleware(next http.Handler, db *sqlx.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Auslesen des "jwt"-Cookies
		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Error(w, "Nicht autorisiert", http.StatusUnauthorized)
			return
		}
		tokenString := cookie.Value

		// Überprüfen des Tokens
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Sicherstellen, dass der Signieralgorithmus passt
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unerwartete Signiermethode: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		// FIRST check if token is valid
		if err != nil || !token.Valid {
			http.Error(w, "Nicht autorisiert", http.StatusUnauthorized)
			return
		}

		// THEN safely access the claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Ungültige Token-Ansprüche", http.StatusUnauthorized)
			return
		}
		userId := cast.ToInt(claims["user_id"])

		newCtx := context.WithValue(r.Context(), "controllerService", controller.CreateService(db, userId))
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}
