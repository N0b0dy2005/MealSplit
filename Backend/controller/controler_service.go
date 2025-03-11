package controller

import (
	"github.com/jmoiron/sqlx"
)

type Service struct {
	DB     *sqlx.DB
	UserID int
}

func CreateService(db *sqlx.DB, userId int) *Service {
	return &Service{
		DB:     db,
		UserID: userId,
	}
}
