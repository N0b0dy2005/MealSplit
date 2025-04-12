package controller

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	DB             *sqlx.DB
	UserID         int
	Admin          int
	ResponseWriter http.ResponseWriter
}

func CreateService(db *sqlx.DB, userId int, adminTest bool, w http.ResponseWriter) *Service {

	// select admin from users where id = ?
	var admin int
	err := db.Get(&admin, "SELECT admin FROM users WHERE id = ?", userId)
	if err != nil {
		log.Println("Error fetching user:", err)
		admin = 0
	}
	return &Service{
		DB:             db,
		UserID:         userId,
		Admin:          admin,
		ResponseWriter: w,
	}
}
