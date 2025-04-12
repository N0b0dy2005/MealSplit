package controller

import (
	"fmt"
	"log"
	"my-backend/graph/model"
	"my-backend/utils/cast"
)

// DebtSummary repräsentiert eine zusammengefasste Schuld
type DebtSummary struct {
	FromUserID  int         `db:"from_user_id"`
	ToUserID    int         `db:"to_user_id"`
	TotalAmount float64     `db:"total_amount"`
	MealsCount  int         `db:"meals_count"`
	CreatedAt   string      `db:"created_at"`
	UpdatedAt   string      `db:"updated_at"`
	IsConfirmed interface{} `db:"is_confirmed"`
}

func (d *Service) GetDebts() ([]*model.Debt, error) {
	fmt.Println("admin ist %v", d.Admin)
	var result []*model.Debt

	// Erstelle eine vereinfachte Abfrage, die Duplikate vermeidet
	var query string
	var args []interface{}

	if d.Admin == 1 {
		// Admin sieht alle Schulden, aber wir filtern Duplikate
		query = `
			SELECT 
				d1.from_user_id, 
				d1.to_user_id, 
				SUM(d1.amount) AS total_amount,
				COUNT(CASE WHEN d1.meals_count > 0 THEN 1 ELSE NULL END) AS meals_count,
				MAX(d1.created_at) AS created_at,
				MAX(d1.updated_at) AS updated_at,
				MAX(d1.is_confirmed) AS is_confirmed
			FROM 
				debts d1
			WHERE
				NOT EXISTS (
					SELECT 1 FROM debts d2 
					WHERE d2.from_user_id = d1.to_user_id 
					AND d2.to_user_id = d1.from_user_id
					AND d1.from_user_id > d1.to_user_id
				)
			GROUP BY 
				d1.from_user_id, d1.to_user_id
			HAVING 
				SUM(d1.amount) > 0
			ORDER BY 
				d1.from_user_id, d1.to_user_id
		`
	} else {
		// Normaler Benutzer sieht nur seine eigenen Schulden
		query = `
			SELECT 
				d1.from_user_id, 
				d1.to_user_id, 
				SUM(d1.amount) AS total_amount,
				COUNT(CASE WHEN d1.meals_count > 0 THEN 1 ELSE NULL END) AS meals_count,
				MAX(d1.created_at) AS created_at,
				MAX(d1.updated_at) AS updated_at,
				MAX(d1.is_confirmed) AS is_confirmed
			FROM 
				debts d1
			WHERE 
				(d1.from_user_id = ? OR d1.to_user_id = ?)
				AND NOT EXISTS (
					SELECT 1 FROM debts d2 
					WHERE d2.from_user_id = d1.to_user_id 
					AND d2.to_user_id = d1.from_user_id
					AND ((d1.from_user_id = ? AND d1.from_user_id > d1.to_user_id) 
						OR (d1.to_user_id = ? AND d1.to_user_id > d1.from_user_id))
				)
			GROUP BY 
				d1.from_user_id, d1.to_user_id
			HAVING 
				SUM(d1.amount) > 0
			ORDER BY 
				CASE 
					WHEN d1.from_user_id = ? THEN 1
					WHEN d1.to_user_id = ? THEN 2
					ELSE 3
				END,
				total_amount DESC
		`
		args = []interface{}{d.UserID, d.UserID, d.UserID, d.UserID, d.UserID, d.UserID}
	}

	var debtSummaries []DebtSummary
	var err error

	if d.Admin == 1 {
		err = d.DB.Select(&debtSummaries, query)
	} else {
		err = d.DB.Select(&debtSummaries, query, args...)
	}

	if err != nil {
		log.Println("Error fetching debts:", err)
		return nil, err
	}

	for _, debt := range debtSummaries {
		isConfirmed := true
		if debt.IsConfirmed != nil {
			isConfirmed = cast.ToBool(debt.IsConfirmed)
		}

		result = append(result, &model.Debt{
			FromUserID:  debt.FromUserID,
			ToUserID:    debt.ToUserID,
			Amount:      cast.ToString(debt.TotalAmount),
			MealsCount:  debt.MealsCount,
			CreateDate:  debt.CreatedAt,
			UpdatedDate: debt.UpdatedAt,
			IsConfirmed: isConfirmed,
		})
	}

	// Einfügen der "Gegeneinträge" für das Frontend
	// Diese Lösung ist spezifisch für dein Frontend-Problem
	if d.Admin != 1 {
		var reverseResults []*model.Debt
		for _, debt := range result {
			if debt.FromUserID == d.UserID {
				// Ich schulde jemandem Geld
				reverseResults = append(reverseResults, debt)
			} else {
				// Jemand schuldet mir Geld
				reverseResults = append(reverseResults, debt)
			}
		}

		result = reverseResults
	}

	return result, nil
}
