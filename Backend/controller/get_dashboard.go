package controller

import (
	"log"
	"my-backend/graph/model"
)

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

	// from_user_id represents the debtor (who owes money)
	// Get top 3 debtors with highest debts
	err = d.DB.Select(&dashboard.tobDebtsPerUser,
		"SELECT from_user_id as userId, SUM(amount) as amount FROM debts GROUP BY from_user_id ORDER BY SUM(amount) DESC LIMIT 3")
	if err != nil {
		log.Println("Error fetching debts per user:", err)
		return nil, err
	}

	// Use total_amount for meals and order by DESC
	err = d.DB.Select(&dashboard.totalCreditsPerUser,
		"SELECT user_id as userId, SUM(total_amount) as amount FROM meals GROUP BY user_id ORDER BY SUM(total_amount) DESC")
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
