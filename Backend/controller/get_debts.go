package controller

import (
	"log"
	"my-backend/graph/model"
	"my-backend/utils/cast"
)

type Debt struct {
	ID int `db:"id"`

	FromUserID int    `db:"from_user_id"`
	ToUserID   int    `db:"to_user_id"`
	Amount     string `db:"amount"`
	MealsCount int    `db:"meals_count"`
	CreatedAt  string `db:"created_at"`
	UpdateAt   string `db:"updated_at"`
}

func (d *Service) GetDebts() ([]*model.Debt, error) {
	var result []*model.Debt

	var debts []Debt
	// Verwende Select statt Get, um mehrere Zeilen abzurufen
	err := d.DB.Select(&debts, "SELECT * FROM debts ORDER BY created_at DESC")
	if err != nil {
		log.Println("Error fetching user:", err)
		return nil, err
	}

	for _, debt := range debts {
		result = append(result, &model.Debt{
			ID:          cast.ToInt(debt.ID),
			Amount:      cast.ToString(debt.Amount),
			FromUserID:  cast.ToInt(debt.FromUserID),
			ToUserID:    cast.ToInt(debt.ToUserID),
			MealsCount:  cast.ToInt(debt.MealsCount),
			CreateDate:  debt.CreatedAt,
			UpdatedDate: debt.UpdateAt,
		})
	}
	return result, nil
}
