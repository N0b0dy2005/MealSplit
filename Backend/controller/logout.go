package controller

import (
	"fmt"
	"my-backend/graph/model"
	"net/http"
	"time"
)

// Logout löscht den JWT-Cookie und meldet den Benutzer ab.
func (s *Service) Logout() (*model.Response, error) {
	fmt.Println("Logout called")

	// Cookie mit abgelaufenem Datum erstellen.
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0), // Datum in der Vergangenheit
		MaxAge:   -1,              // Negative MaxAge löscht den Cookie
		HttpOnly: true,
		// Weitere Attribute (Domain, Secure, SameSite) nur setzen, wenn erforderlich.
	}

	// Cookie im ResponseWriter setzen.
	http.SetCookie(s.ResponseWriter, cookie)

	return &model.Response{
		Status:  "success",
		Message: "Erfolgreich abgemeldet",
		Success: true,
	}, nil
}
