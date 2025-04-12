package controller

import (
	"log"
	"my-backend/graph/model"
	"my-backend/utils/cast"
)

func (d *Service) GetUsers() ([]*model.User, error) {
	var result []*model.User

	var users []User

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
			Admin:       cast.ToBool(user.Admin),
		})
	}

	return result, nil
}

// createfunction GetCurrentUser
func (d *Service) GetCurrentUser() (*model.User, error) {
	var user User
	err := d.DB.Get(&user, "SELECT * FROM users WHERE id = ?", d.UserID)
	if err != nil {
		log.Println("Error fetching user:", err)
		return nil, err
	}

	return &model.User{
		ID:          cast.ToInt(user.ID),
		Name:        cast.ToString(user.Name),
		Email:       cast.ToString(user.Email),
		Debts:       cast.ToString(user.Debts),
		Credits:     cast.ToString(user.Credits),
		CreateDate:  cast.ToString(user.CreateDate),
		UpdatedDate: cast.ToString(user.UpdatedDate),
		PhoneNumber: cast.ToString(user.PhoneNumber),
	}, nil
}
