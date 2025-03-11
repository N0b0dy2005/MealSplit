package controller

import (
	"log"
	"my-backend/graph/model"
	"time"
)

func (d *Service) CreatePost(contend string) (*model.Response, error) {
	// Properly initialize the response object
	response := &model.Response{
		Success: false,
	}

	// get mysql dateTime
	now := time.Now()
	createDate := now.Format("2006-01-02 15:04:05")

	// Note: Ensure the SQL query placeholders match the number of values provided
	_, err := d.DB.Exec(
		`INSERT INTO post(parentid, content, createdate, picture, userid, likes) VALUES (?, ?, ?, ?, ?, ?)`,
		nil, contend, createDate, nil, d.UserID, 0,
	)
	if err != nil {
		log.Println("Error creating post:", err) // Logging the error for debugging
		return response, err
	}

	response.Success = true
	return response, nil
}
