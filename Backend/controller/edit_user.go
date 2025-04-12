package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"my-backend/graph/model"
)

func (d *Service) EditUser(user model.UserInput) (*model.Response, error) {
	// Properly initialize the response object
	response := &model.Response{
		Success: false,
	}

	_, err := d.DB.Exec(`UPDATE users SET name = ?, email = ? WHERE id = ?`, user.Name, user.Email, user.ID)
	if err != nil {
		log.Println("Error deleting user:", err) // Logging the error for debugging
		return response, err
	}

	response.Success = true
	return response, nil
}

func (d *Service) UpdateUser(user model.UpdateUserInput) (*model.Response, error) {
	// Properly initialize the response object
	response := &model.Response{
		Success: false,
	}

	if user.Password == "" {
		fmt.Println("User password is empty, not updating password", user.Password, user.ID, user.Name, user.Email, user.PhoneNumber)
		_, err := d.DB.Exec(`UPDATE users SET name = ?, email = ? ,phone_number = ? WHERE id = ?`, user.Name, user.Email, user.PhoneNumber, user.ID)
		if err != nil {
			log.Println("Error deleting user:", err) // Logging the error for debugging
			return response, err
		}
	} else {
		// Hash das Passwort mit SHA256
		fmt.Println("User password is:", user.Password, user.ID, user.Name, user.Email, user.PhoneNumber)
		hasher := sha256.New()
		hasher.Write([]byte(user.Password))
		passwordHash := hex.EncodeToString(hasher.Sum(nil))

		_, err := d.DB.Exec(`UPDATE users SET name = ?, email = ? ,phone_number = ?, password = ? WHERE id = ?`, user.Name, user.Email, user.PhoneNumber, passwordHash, user.ID)
		if err != nil {
			log.Println("Error updating user:", err) // Logging the error for debugging
			return response, err
		}
	}

	response.Success = true
	return response, nil
}
