package controller

import (
	"log"
	"my-backend/graph/model"
	"time"
)

func (d *Service) DeleteUser(id int) (*model.Response, error) {
	// Properly initialize the response object
	response := &model.Response{
		Success: false,
	}

	// get mysql dateTime
	now := time.Now()
	createDate := now.Format("2006-01-02 15:04:05")
	// Note: Ensure the SQL query placeholders match the number of values provided

	// insert activities
	_, err := d.DB.Exec(`INSERT INTO activities(type,description,from_user_id,to_user_id,user_id, amount,date) VALUES (?,?,?,?,?,?,?)`,
		"user", "User Deleted", nil, nil, nil, 0, createDate)
	if err != nil {
		log.Println("Error creating post:", err) // Logging the error for debugging
		return response, err
	}

	_, err = d.DB.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		log.Println("Error deleting user:", err) // Logging the error for debugging
		return response, err
	}

	response.Success = true
	return response, nil
}
