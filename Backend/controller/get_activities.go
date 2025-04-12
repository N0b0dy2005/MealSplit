package controller

import (
	"log"
	"my-backend/graph/model"
	"my-backend/utils/cast"
)

func (d *Service) GetActivities() ([]*model.Activities, error) {
	var result []*model.Activities

	var activities []Activities
	// Verwende Select statt Get, um mehrere Zeilen abzurufen
	err := d.DB.Select(&activities, "SELECT * FROM activities ORDER BY date ASC LIMIT 5")

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
			IsConfirmed: cast.ToBool(a.IsConfirmed),
		})
	}
	return result, nil
}
