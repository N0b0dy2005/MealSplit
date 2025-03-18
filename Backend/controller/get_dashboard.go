package controller

import (
	"log"
	"my-backend/graph/model"
)

// TotalCreditsPerUser represents the total credits for a user
type TotalCreditsPerUser struct {
	UserId int    `db:"userId"`
	Amount string `db:"amount"`
}

// TobDebtsPerUser represents the total debts for a user
type TobDebtsPerUser struct {
	UserId int    `db:"userId"`
	Amount string `db:"amount"`
}

// Dashboard is the internal representation of dashboard data
type Dashboard struct {
	totalUsers          int
	totalMeals          int
	totalDebts          string
	totalCredits        string
	tobDebtsPerUser     []TobDebtsPerUser
	totalCreditsPerUser []TotalCreditsPerUser
}

// GetDashboard fetches dashboard data from database
func (d *Service) GetDashboard() (*model.Dashboard, error) {
	var dashboard Dashboard

	// Use Get for single-value queries
	err := d.DB.Get(&dashboard.totalUsers, "SELECT Count(*) FROM users")
	if err != nil {
		log.Println("Error fetching user:", err)
		return nil, err
	}

	err = d.DB.Get(&dashboard.totalMeals, "SELECT Count(*) FROM meals")
	if err != nil {
		log.Println("Error fetching meals:", err)
		return nil, err
	}

	// Using the correct column name in debts table
	err = d.DB.Get(&dashboard.totalDebts, "SELECT COALESCE(SUM(amount), '0') FROM debts")
	if err != nil {
		log.Println("Error fetching debts:", err)
		return nil, err
	}

	// Using the correct column name in meals table - total_amount not amount
	err = d.DB.Get(&dashboard.totalCredits, "SELECT COALESCE(SUM(total_amount), '0') FROM meals")
	if err != nil {
		log.Println("Error fetching credits:", err)
		return nil, err
	}

	// Group by from_user_id since that's the debtor
	err = d.DB.Select(&dashboard.tobDebtsPerUser,
		"SELECT from_user_id as userId, SUM(amount) as amount FROM debts GROUP BY from_user_id")
	if err != nil {
		log.Println("Error fetching debts per user:", err)
		return nil, err
	}

	// Use total_amount for meals
	err = d.DB.Select(&dashboard.totalCreditsPerUser,
		"SELECT user_id as userId, SUM(total_amount) as amount FROM meals GROUP BY user_id")
	if err != nil {
		log.Println("Error fetching credits per user:", err)
		return nil, err
	}

	// Convert to model types
	var modelTobDebtsPerUser []*model.TobDebtsPerUser
	for _, debt := range dashboard.tobDebtsPerUser {
		modelTobDebtsPerUser = append(modelTobDebtsPerUser, &model.TobDebtsPerUser{
			UserID: debt.UserId,
			Amount: debt.Amount,
		})
	}

	var modelTotalCreditsPerUser []*model.TotalCreditsPerUser
	for _, credit := range dashboard.totalCreditsPerUser {
		modelTotalCreditsPerUser = append(modelTotalCreditsPerUser, &model.TotalCreditsPerUser{
			UserID: credit.UserId,
			Amount: credit.Amount,
		})
	}

	// order modelTotalCreditsPerUser by amount
	for i := 0; i < len(modelTotalCreditsPerUser); i++ {
		for j := i + 1; j < len(modelTotalCreditsPerUser); j++ {
			if modelTotalCreditsPerUser[i].Amount < modelTotalCreditsPerUser[j].Amount {
				modelTotalCreditsPerUser[i], modelTotalCreditsPerUser[j] = modelTotalCreditsPerUser[j], modelTotalCreditsPerUser[i]
			}
		}
	}

	// order modelTobDebtsPerUser by amount
	for i := 0; i < len(modelTobDebtsPerUser); i++ {
		for j := i + 1; j < len(modelTobDebtsPerUser); j++ {
			if modelTobDebtsPerUser[i].Amount < modelTobDebtsPerUser[j].Amount {
				modelTobDebtsPerUser[i], modelTobDebtsPerUser[j] = modelTobDebtsPerUser[j], modelTobDebtsPerUser[i]
			}
		}
	}

	result := &model.Dashboard{
		TotalUsers:          dashboard.totalUsers,
		TotalMeals:          dashboard.totalMeals,
		TotalDebts:          dashboard.totalDebts,
		TotalCredits:        dashboard.totalCredits,
		TobDebtsPerUser:     modelTobDebtsPerUser,
		TotalCreditsPerUser: modelTotalCreditsPerUser,
	}
	return result, nil
}
