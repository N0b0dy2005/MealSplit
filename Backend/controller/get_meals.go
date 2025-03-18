package controller

import (
	"log"
	"my-backend/graph/model"
	"my-backend/utils/cast"
)

type Meals struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Date        string `db:"date"`
	TotalAmount string `db:"total_amount"`
	UserID      int    `db:"user_id"` // Updated tag
	Description string `db:"description"`
	CreateDate  string `db:"created_at"`
	UpdatedDate string `db:"updated_at"`
	UserIds     string `db:"user_ids"`
}

func (d *Service) GetMeals() ([]*model.Meal, error) {
	var result []*model.Meal

	var meals []Meals
	// Verwende Select statt Get, um mehrere Zeilen abzurufen
	err := d.DB.Select(&meals, "SELECT * FROM meals ORDER BY created_at DESC")
	if err != nil {
		log.Println("Error fetching user:", err)
		return nil, err
	}

	for _, meal := range meals {
		result = append(result, &model.Meal{
			ID:          cast.ToInt(meal.ID),
			Name:        meal.Name,
			Date:        meal.Date,
			TotalAmount: meal.TotalAmount,
			UserID:      cast.ToInt(meal.UserID),
			Description: meal.Description,
			CreateDate:  meal.CreateDate,
			UpdatedDate: meal.UpdatedDate,
			UserIds:     meal.UserIds,
		})
	}
	return result, nil
}
