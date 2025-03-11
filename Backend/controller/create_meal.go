package controller

import (
	"database/sql"
	"log"
	"my-backend/graph/model"
	"my-backend/utils/cast"
	"strconv"
	"strings"
	"time"
)

type Debts struct {
	ID         int     `db:"id"`
	FromUserID int     `db:"from_user_id"`
	ToUserID   int     `db:"to_user_id"`
	Amount     float64 `db:"amount"`
	MealCount  int     `db:"meals_count"`
	CreatedAt  string  `db:"created_at"`
	UpdatedAt  string  `db:"updated_at"`
}

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

	// Insert meal record
	_, err = d.DB.Exec(
		`INSERT INTO meals(name,date,total_amount,user_id,description,created_at,updated_at,user_ids) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		meal.Name, meal.Date, meal.TotalAmount, meal.UserID, meal.Description, createDate, createDate, meal.UserIds)

	if err != nil {
		log.Println("Error creating meal:", err)
		return response, err
	}

	// Convert user IDs from string to array
	userIds, err := GetIntArrayFromString(meal.UserIds)
	if err != nil {
		log.Println("Error converting user ids:", err)
		return response, err
	}

	// Calculate per-person share
	totalAmount := cast.ToFloat64(meal.TotalAmount)
	perPersonAmount := totalAmount / float64(len(userIds))

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

		// Check for existing debt FROM userId TO meal.UserID
		var debtFromUserToPayer Debts
		err = d.DB.Get(&debtFromUserToPayer,
			"SELECT * FROM debts WHERE from_user_id = ? AND to_user_id = ?",
			userId, meal.UserID)

		// Check for error but don't exit if it's just "no rows"
		hasDebtFromUserToPayer := true
		if err != nil {
			if err == sql.ErrNoRows {
				hasDebtFromUserToPayer = false
			} else {
				log.Println("Database error checking debts:", err)
				return response, err
			}
		}

		// Check for existing debt FROM meal.UserID TO userId (reverse)
		var debtFromPayerToUser Debts
		err = d.DB.Get(&debtFromPayerToUser,
			"SELECT * FROM debts WHERE from_user_id = ? AND to_user_id = ?",
			meal.UserID, userId)

		// Check for error but don't exit if it's just "no rows"
		hasDebtFromPayerToUser := true
		if err != nil {
			if err == sql.ErrNoRows {
				hasDebtFromPayerToUser = false
			} else {
				log.Println("Database error checking reverse debts:", err)
				return response, err
			}
		}

		// Process debts based on what exists
		if hasDebtFromUserToPayer {
			// User already owes the payer - increase the debt
			newAmount := debtFromUserToPayer.Amount + perPersonAmount
			newMealCount := debtFromUserToPayer.MealCount + 1

			_, err = d.DB.Exec(
				`UPDATE debts SET amount = ?, meals_count = ?, updated_at = ? WHERE id = ?`,
				newAmount, newMealCount, createDate, debtFromUserToPayer.ID)

			if err != nil {
				log.Println("Error updating existing debt:", err)
				return response, err
			}
		} else if hasDebtFromPayerToUser {
			// Payer owes the user - reduce or reverse this debt
			existingAmount := debtFromPayerToUser.Amount

			if existingAmount > perPersonAmount {
				// Just reduce the debt
				newAmount := existingAmount - perPersonAmount
				_, err = d.DB.Exec(
					`UPDATE debts SET amount = ?, updated_at = ? WHERE id = ?`,
					newAmount, createDate, debtFromPayerToUser.ID)
			} else if existingAmount < perPersonAmount {
				// Reverse the debt direction
				newAmount := perPersonAmount - existingAmount

				// Delete the old debt record
				_, err = d.DB.Exec(`DELETE FROM debts WHERE id = ?`, debtFromPayerToUser.ID)
				if err != nil {
					log.Println("Error deleting old debt:", err)
					return response, err
				}

				// Create new debt in opposite direction
				_, err = d.DB.Exec(
					`INSERT INTO debts(from_user_id, to_user_id, amount, meals_count, created_at, updated_at) 
					 VALUES (?, ?, ?, ?, ?, ?)`,
					userId, meal.UserID, newAmount, 1, createDate, createDate)
			} else {
				// Debts cancel out exactly - delete the record
				_, err = d.DB.Exec(`DELETE FROM debts WHERE id = ?`, debtFromPayerToUser.ID)
			}

			if err != nil {
				log.Println("Error processing debt update:", err)
				return response, err
			}
		} else {
			// No existing debt in either direction - create new debt
			_, err = d.DB.Exec(
				`INSERT INTO debts(from_user_id, to_user_id, amount, meals_count, created_at, updated_at) 
				 VALUES (?, ?, ?, ?, ?, ?)`,
				userId, meal.UserID, perPersonAmount, 1, createDate, createDate)

			if err != nil {
				log.Println("Error creating new debt:", err)
				return response, err
			}
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

// Neue Funktion zur Aktualisierung der Benutzer-Schulden und Guthaben
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
