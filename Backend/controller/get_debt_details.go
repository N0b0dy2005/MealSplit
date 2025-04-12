package controller

import (
	"log"
	"my-backend/graph/model"
	"my-backend/utils/cast"
)

// DebtDetail repräsentiert einen einzelnen Eintrag in der Debts-Tabelle

// GetDebtDetails gibt alle Einträge zurück, die zu einer Schuldenbeziehung zwischen zwei Benutzern gehören
func (d *Service) GetDebtDetails(fromUserID, toUserID int) ([]*model.DebtDetail, error) {
	var debtDetails []DebtDetail
	var result []*model.DebtDetail

	// Hole alle Einträge zwischen den beiden Benutzern in beide Richtungen
	query := `
		SELECT * FROM debts 
		WHERE (from_user_id = ? AND to_user_id = ?) 
		   OR (from_user_id = ? AND to_user_id = ?)
		ORDER BY created_at DESC
	`

	err := d.DB.Select(&debtDetails, query, fromUserID, toUserID, toUserID, fromUserID)
	if err != nil {
		log.Println("Error fetching debt details:", err)
		return nil, err
	}

	// Transformiere die DB-Entitäten in das Response-Format
	for _, detail := range debtDetails {
		isConfirmed := true
		if detail.IsConfirmed != nil {
			isConfirmed = cast.ToBool(detail.IsConfirmed)
		}

		var mealID, paymentID *int
		if detail.MealID != nil {
			mealIDCopy := *detail.MealID
			mealID = &mealIDCopy
		}

		// Bestimme den Typ des Eintrags
		entryType := "unknown"
		if detail.MealID != nil {
			entryType = "meal"
		} else if detail.PaymentID != nil {
			entryType = "payment"
		}

		// Hole zusätzliche Informationen für Mahlzeiten, falls verfügbar
		var mealInfo *model.MealInfo
		if detail.MealID != nil {
			var meal MealInfo
			err := d.DB.Get(&meal, "SELECT * FROM meals WHERE id = ?", *detail.MealID)
			if err == nil {
				description := meal.Description
				mealInfo = &model.MealInfo{
					ID:          meal.ID,
					Name:        meal.Name,
					Description: &description,
					Date:        meal.Date,
					TotalAmount: cast.ToString(meal.TotalAmount),
				}
			} else {
				log.Printf("Fehler beim Abrufen der Mahlzeitinformationen für ID %d: %v", *detail.MealID, err)
			}
		}

		// Hole zusätzliche Informationen für Zahlungen, falls verfügbar
		var paymentInfo *model.PaymentInfo
		if detail.PaymentID != nil {
			paymentIDCopy := *detail.PaymentID
			paymentID = &paymentIDCopy

			var payment PaymentInfo
			err := d.DB.Get(&payment, "SELECT * FROM payment WHERE id = ?", *detail.PaymentID)
			if err == nil {
				description := payment.Description
				paymentInfo = &model.PaymentInfo{
					ID:          payment.ID,
					Description: &description,
					Date:        payment.Date,
					Amount:      payment.Amount,
				}
			} else {
				log.Printf("Fehler beim Abrufen der Zahlungsinformationen für ID %d: %v", *detail.PaymentID, err)
			}
		}

		// Benutzer-Informationen hinzufügen
		var fromUser, toUser string
		d.DB.Get(&fromUser, "SELECT name FROM users WHERE id = ?", detail.FromUserID)
		d.DB.Get(&toUser, "SELECT name FROM users WHERE id = ?", detail.ToUserID)

		result = append(result, &model.DebtDetail{
			ID:           detail.ID,
			FromUserID:   detail.FromUserID,
			ToUserID:     detail.ToUserID,
			FromUserName: fromUser,
			ToUserName:   toUser,
			Amount:       cast.ToString(detail.Amount),
			MealsCount:   detail.MealsCount,
			CreateDate:   detail.CreatedAt,
			IsConfirmed:  isConfirmed,
			Type:         entryType,
			MealID:       mealID,
			PaymentID:    paymentID,
			MealInfo:     mealInfo,
			PaymentInfo:  paymentInfo,
		})
	}

	return result, nil
}
