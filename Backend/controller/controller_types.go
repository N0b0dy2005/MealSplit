package controller

import (
	"database/sql"
)

// DebtDetail repräsentiert einen einzelnen Eintrag in der Debts-Historie
type DebtDetail struct {
	ID          int         `db:"id"`
	FromUserID  int         `db:"from_user_id"`
	ToUserID    int         `db:"to_user_id"`
	Amount      float64     `db:"amount"`
	MealsCount  int         `db:"meals_count"`
	CreatedAt   string      `db:"created_at"`
	UpdatedAt   string      `db:"updated_at"`
	IsConfirmed interface{} `db:"is_confirmed"`
	MealID      *int        `db:"meal_id"`
	PaymentID   *int        `db:"payment_id"`
}

// MealInfo repräsentiert eine Mahlzeit aus der meals-Tabelle
type MealInfo struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	Date        string  `db:"date"`
	TotalAmount float64 `db:"total_amount"`
	UserID      int     `db:"user_id"`
	Description string  `db:"description"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   string  `db:"updated_at"`
	UserIds     string  `db:"user_ids"`
	Products    string  `db:"products"`
}

// PaymentInfo repräsentiert eine Zahlung aus der payment-Tabelle
type PaymentInfo struct {
	ID          int    `db:"id"`
	FromUserID  int    `db:"from_user_id"`
	ToUserID    int    `db:"to_user_id"`
	Amount      string `db:"amount"`
	Description string `db:"description"`
	Date        string `db:"date"`
	IsConfirmed int    `db:"is_confirmed"`
}

// Debt struct aus get_debts.go
type Debt struct {
	ID          int           `db:"id"`
	FromUserID  int           `db:"from_user_id"`
	ToUserID    int           `db:"to_user_id"`
	Amount      string        `db:"amount"`
	MealsCount  int           `db:"meals_count"`
	CreatedAt   string        `db:"created_at"`
	UpdateAt    string        `db:"updated_at"`
	IsConfirmed sql.NullInt64 `db:"is_confirmed"`
	MealID      int           `db:"meal_id"`
}

// Product struct aus create_meal.go
type Product struct {
	Name                 string  `json:"name"`
	Price                float64 `json:"price"`
	IsSpecific           bool    `json:"isSpecific"`
	SpecificParticipants []int   `json:"specificParticipants"`
}

// Meals struct aus get_meals.go
type Meals struct {
	ID          int            `db:"id"`
	Name        string         `db:"name"`
	Date        string         `db:"date"`
	TotalAmount string         `db:"total_amount"`
	UserID      int            `db:"user_id"`
	Description string         `db:"description"`
	CreateDate  string         `db:"created_at"`
	UpdatedDate string         `db:"updated_at"`
	UserIds     string         `db:"user_ids"`
	Products    sql.NullString `db:"products"`
}

// Payment struct aus get_payment.go
type Payment struct {
	ID          int           `db:"id"`
	FromUserID  int           `db:"from_user_id"`
	ToUserID    int           `db:"to_user_id"`
	Amount      string        `db:"amount"`
	Description string        `db:"description"`
	Date        string        `db:"date"`
	IsConfirmed sql.NullInt64 `db:"is_confirmed"`
}

// User struct aus get_users.go
type User struct {
	ID          int            `db:"id"`
	Name        sql.NullString `db:"name"`
	Email       sql.NullString `db:"email"`
	Debts       sql.NullString `db:"debts"`
	Credits     sql.NullString `db:"credits"`
	CreateDate  sql.NullString `db:"created_at"`
	UpdatedDate sql.NullString `db:"updated_at"`
	Password    sql.NullString `db:"password"`
	PhoneNumber sql.NullString `db:"phone_number"`
	Admin       sql.NullString `db:"admin"`
}

// TotalCreditsPerUser aus get_dashboard.go
type TotalCreditsPerUser struct {
	UserId int    `db:"userId"`
	Amount string `db:"amount"`
}

// TobDebtsPerUser aus get_dashboard.go
type TobDebtsPerUser struct {
	UserId int    `db:"userId"`
	Amount string `db:"amount"`
}

// Dashboard aus get_dashboard.go
type Dashboard struct {
	totalUsers          int
	totalMeals          int
	totalDebts          string
	totalCredits        string
	tobDebtsPerUser     []TobDebtsPerUser
	totalCreditsPerUser []TotalCreditsPerUser
}

// Activities aus get_activities.go
type Activities struct {
	ID          int           `db:"id"`
	Type        string        `db:"type"`
	Description string        `db:"description"`
	FromUserID  sql.NullInt64 `db:"from_user_id"`
	ToUserID    sql.NullInt64 `db:"to_user_id"`
	UserID      sql.NullInt64 `db:"user_id"`
	Amount      string        `db:"amount"`
	Date        string        `db:"date"`
	IsConfirmed sql.NullInt64 `db:"is_confirmed"`
}

// item aus upload_receipt.go
type item struct {
	Name  string
	Price float64
}
