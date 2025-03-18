package controller

import (
	"database/sql"
	"log"
	"my-backend/graph/model"
	"my-backend/utils/cast"
)

type User struct {
	ID          int            `db:"id"`
	Name        sql.NullString `db:"name"`
	Email       sql.NullString `db:"email"`
	Debts       sql.NullString `db:"debts"`
	Credits     sql.NullString `db:"credits"`
	CreateDate  sql.NullString `db:"created_at"` // angepasst!
	UpdatedDate sql.NullString `db:"updated_at"`
	Password    sql.NullString `db:"password"`
	PhoneNumber sql.NullString `db:"phone_number"`
}

func (d *Service) GetUsers() ([]*model.User, error) {
	var result []*model.User

	var users []User
	// Verwende Select statt Get, um mehrere Zeilen abzurufen
	err := d.DB.Select(&users, "SELECT * FROM users ORDER BY name ASC ")
	if err != nil {
		log.Println("Error fetching user:", err)
		return nil, err
	}

	for _, user := range users {
		result = append(result, &model.User{
			ID:          cast.ToInt(user.ID),
			Name:        cast.ToString(user.Name),
			Email:       cast.ToString(user.Email),
			Debts:       cast.ToString(user.Debts),
			Credits:     cast.ToString(user.Credits),
			CreateDate:  cast.ToString(user.CreateDate),
			UpdatedDate: cast.ToString(user.UpdatedDate),
			PhoneNumber: cast.ToString(user.PhoneNumber),
		})
	}

	return result, nil
}
