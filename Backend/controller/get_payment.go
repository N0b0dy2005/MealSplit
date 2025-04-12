package controller

import (
	"log"
	"my-backend/graph/model"
	"my-backend/utils/cast"
)

func (d *Service) GetPayment() ([]*model.Payment, error) {
	var result []*model.Payment

	var payments []Payment
	// Verwende Select statt Get, um mehrere Zeilen abzurufen
	var err error

	if d.Admin == 1 {
		// Wenn der Benutzer ein Admin ist, zeige alle Zahlungen
		err = d.DB.Select(&payments, "SELECT * FROM payment ORDER BY date DESC LIMIT 5")
	} else {
		// Wenn der Benutzer kein Admin ist, zeige nur die Zahlungen, die mit dem Benutzer zusammenhängen
		err = d.DB.Select(&payments, "SELECT * FROM payment WHERE from_user_id = ? OR to_user_id = ? ORDER BY date DESC LIMIT 5", d.UserID, d.UserID)
	}

	if err != nil {
		log.Println("Error fetching user:", err)
		return nil, err
	}

	var isConfirmed bool

	for _, payment := range payments {
		if payment.IsConfirmed.Valid {
			isConfirmed = cast.ToBool(payment.IsConfirmed)
		} else {
			isConfirmed = true
		}
		result = append(result, &model.Payment{
			ID:          cast.ToInt(payment.ID),
			Amount:      cast.ToString(payment.Amount),
			FromUserID:  cast.ToInt(payment.FromUserID),
			ToUserID:    cast.ToInt(payment.ToUserID),
			Description: cast.ToString(payment.Description),
			Date:        cast.ToString(payment.Date),
			IsConfirmed: isConfirmed,
		})
	}
	return result, nil
}
