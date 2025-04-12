package controller

import (
	"encoding/json"
	"log"
	"my-backend/graph/model"
	"my-backend/utils/cast"
	"strconv"
	"strings"
	"time"
)

func (d *Service) CreateMeal(meal model.MealInput) (*model.Response, error) {
	response := &model.Response{
		Success: false,
	}

	// MySQL datetime format
	now := time.Now()
	createDate := now.Format("2006-01-02 15:04:05")

	// insert activities
	_, err := d.DB.Exec(`INSERT INTO activities(type,description,from_user_id,to_user_id,user_id, amount,date) VALUES (?,?,?,?,?,?,?)`,
		"meal", meal.Description, nil, nil, meal.UserID, meal.TotalAmount, meal.Date)
	if err != nil {
		log.Println("Error creating post:", err) // Logging the error for debugging
		return response, err
	}

	// Insert meal record mit Produkten
	result, err := d.DB.Exec(
		`INSERT INTO meals(name,date,total_amount,user_id,description,created_at,updated_at,user_ids,products) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		meal.Name, meal.Date, meal.TotalAmount, meal.UserID, meal.Description, createDate, createDate, meal.UserIds, meal.Produkts)

	if err != nil {
		log.Println("Error creating meal:", err)
		return response, err
	}

	// Get the newly created meal ID
	mealID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting meal ID:", err)
		return response, err
	}

	// Convert user IDs from string to array
	userIds, err := GetIntArrayFromString(meal.UserIds)
	if err != nil {
		log.Println("Error converting user ids:", err)
		return response, err
	}

	// Verarbeite die Produkte, wenn vorhanden
	specificCosts := make(map[int]float64)
	totalSpecificCosts := 0.0

	if meal.Produkts != "" {
		var products []Product
		err = json.Unmarshal([]byte(meal.Produkts), &products)
		if err != nil {
			log.Println("Error parsing products:", err)
			return response, err
		}

		// Berechne spezifische Kosten pro Benutzer
		for _, product := range products {
			if product.IsSpecific && len(product.SpecificParticipants) > 0 {
				// Teile den Preis durch die Anzahl der spezifischen Teilnehmer
				pricePerParticipant := product.Price / float64(len(product.SpecificParticipants))
				totalSpecificCosts += product.Price

				for _, userId := range product.SpecificParticipants {
					specificCosts[userId] += pricePerParticipant
				}
			}
		}
	}

	// Berechne den zu teilenden Betrag (Gesamtbetrag - spezifische Kosten)
	totalAmount := cast.ToFloat64(meal.TotalAmount)
	sharedAmount := totalAmount - totalSpecificCosts

	// Erstelle eine Liste der Benutzer, die am geteilten Betrag teilnehmen
	sharedUsers := make(map[int]bool)
	for _, userId := range userIds {
		sharedUsers[userId] = true
	}

	// Berechne den Anteil pro Person für den zu teilenden Betrag
	perPersonSharedAmount := 0.0
	if len(sharedUsers) > 0 {
		perPersonSharedAmount = sharedAmount / float64(len(sharedUsers))
	}

	// Halte eine Liste der Benutzer, die aktualisiert werden müssen
	usersToUpdate := make(map[int]bool)
	usersToUpdate[meal.UserID] = true // Der Zahler muss auf jeden Fall aktualisiert werden

	// Process each user's debt
	for _, userId := range userIds {
		// Skip the person who paid
		if userId == meal.UserID {
			continue
		}

		// Füge den Benutzer zur Liste der zu aktualisierenden Benutzer hinzu
		usersToUpdate[userId] = true

		// Berechne Gesamtbetrag für diesen Benutzer (geteilter Betrag + spezifische Kosten)
		userTotalAmount := perPersonSharedAmount + specificCosts[userId]

		// Statt zu prüfen, ob Schulden existieren und zu aktualisieren,
		// erstellen wir immer einen neuen Eintrag für diese Mahlzeit
		_, err = d.DB.Exec(
			`INSERT INTO debts(from_user_id, to_user_id, amount, meals_count, created_at, updated_at, meal_id) 
				VALUES (?, ?, ?, ?, ?, ?, ?)`,
			userId, meal.UserID, userTotalAmount, 1, createDate, createDate, mealID)

		if err != nil {
			log.Println("Error creating new debt:", err)
			return response, err
		}
	}

	// Aktualisiere alle betroffenen Benutzer in der users-Tabelle
	for userId := range usersToUpdate {
		err = d.UpdateUserDebtsAndCredits(userId)
		if err != nil {
			log.Println("Error updating user debt information:", err)
			return response, err
		}
	}

	response.Success = true
	return response, nil
}

// Funktion zur Aktualisierung der Benutzer-Schulden und Guthaben
func (d *Service) UpdateUserDebtsAndCredits(userId int) error {
	// Berechne Gesamtschulden (was der Benutzer anderen schuldet)
	var totalDebts float64
	err := d.DB.Get(&totalDebts, `
		SELECT COALESCE(SUM(amount), 0) 
		FROM debts 
		WHERE from_user_id = ?`, userId)
	if err != nil {
		return err
	}

	// Berechne Gesamtguthaben (was andere dem Benutzer schulden)
	var totalCredits float64
	err = d.DB.Get(&totalCredits, `
		SELECT COALESCE(SUM(amount), 0) 
		FROM debts 
		WHERE to_user_id = ?`, userId)
	if err != nil {
		return err
	}

	// Aktualisiere die users-Tabelle
	_, err = d.DB.Exec(`
		UPDATE users 
		SET debts = ?, credits = ? 
		WHERE id = ?`, totalDebts, totalCredits, userId)
	if err != nil {
		return err
	}

	return nil
}

func GetIntArrayFromString(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	var result []int

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		num, err := strconv.Atoi(trimmed)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}
