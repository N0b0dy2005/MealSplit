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
	err := d.DB.Select(&payments, "SELECT * FROM payment")
	if err != nil {
		log.Println("Error fetching user:", err)
		return nil, err
	}

	for _, payment := range payments {
		result = append(result, &model.Payment{
			ID:          cast.ToInt(payment.ID),
			Amount:      cast.ToString(payment.Amount),
			FromUserID:  cast.ToInt(payment.FromUserID),
			ToUserID:    cast.ToInt(payment.ToUserID),
			Description: cast.ToString(payment.Description),
			Date:        cast.ToString(payment.Date),
		})
	}
	return result, nil
}
