package controller

import (
	"log"
	"my-backend/graph/model"
	"my-backend/utils/cast"
)

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
			Products:    cast.ToString(meal.Products),
		})
	}
	return result, nil
}
func (d *Service) GetMealByID(id int) (*model.Meal, error) {
	var result *model.Meal

	var meal Meals
	// Verwende Select statt Get, um mehrere Zeilen abzurufen
	err := d.DB.Select(&meal, "SELECT * FROM meals WHERE id = ? ORDER BY created_at DESC", id)
	if err != nil {
		log.Println("Error fetching user:", err)
		return nil, err
	}

	result = &model.Meal{
		ID:          cast.ToInt(meal.ID),
		Name:        meal.Name,
		Date:        meal.Date,
		TotalAmount: meal.TotalAmount,
		UserID:      cast.ToInt(meal.UserID),
		Description: meal.Description,
		CreateDate:  meal.CreateDate,
		UpdatedDate: meal.UpdatedDate,
		UserIds:     meal.UserIds,
		Products:    cast.ToString(meal.Products),
	}

	return result, nil
}

func (d *Service) GetMealsForCurrentUser() ([]*model.Meal, error) {
	var result []*model.Meal

	var meals []Meals
	// Verwende Select statt Get, um mehrere Zeilen abzurufen
	err := d.DB.Select(&meals, "SELECT * FROM meals WHERE user_id = ? ORDER BY created_at DESC", d.UserID)
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
			Products:    cast.ToString(meal.Products),
		})
	}

	return result, nil
}
