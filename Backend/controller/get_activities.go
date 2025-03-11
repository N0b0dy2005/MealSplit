package controller

import (
	"database/sql"
	"log"
	"my-backend/graph/model"
	"my-backend/utils/cast"
)

type Activities struct {
	ID          int           `db:"id"`
	Type        string        `db:"type"`
	Description string        `db:"description"`
	FromUserID  sql.NullInt64 `db:"from_user_id"`
	ToUserID    sql.NullInt64 `db:"to_user_id"`
	UserID      sql.NullInt64 `db:"user_id"`
	Amount      string        `db:"amount"`
	Date        string        `db:"date"`
}

func (d *Service) GetActivities() ([]*model.Activities, error) {
	var result []*model.Activities

	var activities []Activities
	// Verwende Select statt Get, um mehrere Zeilen abzurufen
	err := d.DB.Select(&activities, "SELECT * FROM activities")
	if err != nil {
		log.Println("Error fetching user:", err)
		return nil, err
	}

	for _, a := range activities {
		result = append(result, &model.Activities{
			ID:          a.ID,
			Type:        a.Type,
			Description: cast.ToString(a.Description),
			FromUserID:  cast.ToInt(a.FromUserID),
			ToUserID:    cast.ToInt(a.ToUserID),
			UserID:      cast.ToInt(a.UserID),
			Amount:      a.Amount,
			Date:        a.Date,
		})
	}
	return result, nil
}
