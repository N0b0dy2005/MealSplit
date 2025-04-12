package controller

import (
	"fmt"
	"log"
	"my-backend/graph/model"
	"my-backend/utils/cast"
	"time"
)

func (d *Service) CreatePayment(payment model.PaymentInput) (*model.Response, error) {
	response := &model.Response{
		Success: false,
	}
	// MySQL datetime format
	now := time.Now()
	createDate := now.Format("2006-01-02 15:04:05")

	is_confirmed := 1
	if payment.FromUserID == d.UserID && d.Admin == 0 {
		is_confirmed = 0
	}
	// insert activities
	_, err := d.DB.Exec(`INSERT INTO activities(type,description,from_user_id,to_user_id,user_id, amount,date,is_confirmed) VALUES (?,?,?,?,?,?,?,?)`,
		"payment", payment.Description, payment.FromUserID, payment.ToUserID, payment.FromUserID, payment.Amount, createDate, is_confirmed)
	if err != nil {
		log.Println("Error creating post:", err) // Logging the error for debugging
		return response, err
	}

	// Debug-Ausgabe
	fmt.Println("Debug: Erstelle Zahlung mit FromUserID:", payment.FromUserID,
		"ToUserID:", payment.ToUserID, "Amount:", payment.Amount)

	// Insert new payment record
	result, err := d.DB.Exec(
		`INSERT INTO payment(from_user_id, to_user_id, amount, description, date, is_confirmed) 
     VALUES (?, ?, ?, ?, ?, ?)`,
		payment.FromUserID, payment.ToUserID, payment.Amount, payment.Description,
		createDate, is_confirmed)

	if err != nil {
		log.Println("Debug: Error creating payment:", err)
		return response, err
	} else {
		log.Println("Debug: Successfully inserted new payment record")
	}

	// Get the payment ID
	paymentID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting payment ID:", err)
		return response, err
	}

	// Verarbeite die Zahlung und erstelle einen Eintrag in der Debts-Tabelle
	err = d.ProcessConfirmedPayment(payment, createDate, cast.ToBool(is_confirmed), paymentID)
	if err != nil {
		return response, err
	}

	response.Success = true
	return response, nil
}

// ProcessConfirmedPayment handles debt creation for payments
func (d *Service) ProcessConfirmedPayment(payment model.PaymentInput, createDate string, isConfirmed bool, paymentID int64) error {
	// Bei einer bestätigten Zahlung einen neuen Eintrag in der Debts-Tabelle erstellen
	if isConfirmed {
		// Erstelle einen neuen Eintrag - der Empfänger schuldet dem Sender jetzt Geld
		_, err := d.DB.Exec(
			`INSERT INTO debts(from_user_id, to_user_id, amount, meals_count, created_at, updated_at, payment_id) 
             VALUES (?, ?, ?, ?, ?, ?, ?)`,
			payment.ToUserID, payment.FromUserID, payment.Amount, 0, createDate, createDate, paymentID)

		if err != nil {
			log.Println("Error creating debt:", err)
			return err
		}
	}

	// Aktualisiere die Benutzerinformationen für beide beteiligten Benutzer
	err := d.UpdateUserDebtsAndCredits(payment.FromUserID)
	if err != nil {
		log.Println("Error updating sender's debt information:", err)
		return err
	}

	err = d.UpdateUserDebtsAndCredits(payment.ToUserID)
	if err != nil {
		log.Println("Error updating receiver's debt information:", err)
		return err
	}

	return nil
}
