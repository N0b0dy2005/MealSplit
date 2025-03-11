package controller

import (
	"database/sql"
	"log"
	"my-backend/graph/model"
	"my-backend/utils/cast"
	"time"
)

type Payment struct {
	ID          int    `db:"id"`
	FromUserID  int    `db:"from_user_id"`
	ToUserID    int    `db:"to_user_id"`
	Amount      string `db:"amount"`
	Description string `db:"description"`
	Date        string `db:"date"`
}

func (d *Service) CreatePayment(payment model.PaymentInput) (*model.Response, error) {
	response := &model.Response{
		Success: false,
	}
	// MySQL datetime format
	now := time.Now()
	createDate := now.Format("2006-01-02 15:04:05")

	// insert activities
	_, err := d.DB.Exec(`INSERT INTO activities(type,description,from_user_id,to_user_id,user_id, amount,date) VALUES (?,?,?,?,?,?,?)`,
		"payment", payment.Description, payment.FromUserID, payment.ToUserID, payment.FromUserID, payment.Amount, createDate)
	if err != nil {
		log.Println("Error creating post:", err) // Logging the error for debugging
		return response, err
	}

	// Speichere die Zahlung in der payment-Tabelle
	_, err = d.DB.Exec(
		`INSERT INTO payment(from_user_id, to_user_id, amount, description, date) 
         VALUES (?, ?, ?, ?, ?)`,
		payment.FromUserID, payment.ToUserID, payment.Amount, payment.Description, createDate)

	if err != nil {
		log.Println("Error creating payment:", err)
		return response, err
	}

	// Prüfe, ob der Sender (FromUserID) bereits Schulden beim Empfänger (ToUserID) hat
	var debtFromSenderToReceiver Debts
	err = d.DB.Get(&debtFromSenderToReceiver,
		"SELECT * FROM debts WHERE from_user_id = ? AND to_user_id = ?",
		payment.FromUserID, payment.ToUserID)

	// Prüfe, ob error "no rows" ist oder ein anderer Fehler
	hasDebtFromSenderToReceiver := true
	if err != nil {
		if err == sql.ErrNoRows {
			hasDebtFromSenderToReceiver = false
		} else {
			log.Println("Database error checking debts:", err)
			return response, err
		}
	}

	// Prüfe, ob der Empfänger (ToUserID) bereits Schulden beim Sender (FromUserID) hat
	var debtToFromSender Debts
	err = d.DB.Get(&debtToFromSender,
		"SELECT * FROM debts WHERE from_user_id = ? AND to_user_id = ?",
		payment.ToUserID, payment.FromUserID)

	// Prüfe, ob error "no rows" ist oder ein anderer Fehler
	hasDebtToFromSender := true
	if err != nil {
		if err == sql.ErrNoRows {
			hasDebtToFromSender = false
		} else {
			log.Println("Database error checking debts:", err)
			return response, err
		}
	}

	// Aktualisiere die Schuldentabelle basierend auf der Zahlung
	if hasDebtFromSenderToReceiver {
		// Der Sender schuldet dem Empfänger bereits Geld - Reduziere diese Schuld
		newAmount := cast.ToFloat64(debtFromSenderToReceiver.Amount) - cast.ToFloat64(payment.Amount)

		if newAmount > 0 {
			// Reduziere die bestehende Schuld
			_, err = d.DB.Exec(
				`UPDATE debts SET amount = ?, updated_at = ? WHERE id = ?`,
				newAmount, createDate, debtFromSenderToReceiver.ID)
		} else if newAmount < 0 {
			// Lösche die alte Schuld und erstelle eine neue in umgekehrter Richtung
			_, err = d.DB.Exec(`DELETE FROM debts WHERE id = ?`, debtFromSenderToReceiver.ID)
			if err == nil {
				_, err = d.DB.Exec(
					`INSERT INTO debts(from_user_id, to_user_id, amount, meals_count, created_at, updated_at) 
                     VALUES (?, ?, ?, ?, ?, ?)`,
					payment.ToUserID, payment.FromUserID, -newAmount, 0, createDate, createDate)
			}
		} else { // newAmount == 0
			// Schulden sind vollständig beglichen - Lösche den Eintrag
			_, err = d.DB.Exec(`DELETE FROM debts WHERE id = ?`, debtFromSenderToReceiver.ID)
		}

		if err != nil {
			log.Println("Error updating debts:", err)
			return response, err
		}
	} else if hasDebtToFromSender {
		// Der Empfänger schuldet dem Sender bereits Geld - Erhöhe diese Schuld
		newAmount := cast.ToFloat64(debtToFromSender.Amount) + cast.ToFloat64(payment.Amount)

		_, err = d.DB.Exec(
			`UPDATE debts SET amount = ?, updated_at = ? WHERE id = ?`,
			newAmount, createDate, debtToFromSender.ID)

		if err != nil {
			log.Println("Error updating debts:", err)
			return response, err
		}
	} else {
		// Keine bestehende Schuld in beide Richtungen
		// Der Empfänger schuldet nun dem Sender Geld (weil Sender Geld überwiesen hat)
		_, err = d.DB.Exec(
			`INSERT INTO debts(from_user_id, to_user_id, amount, meals_count, created_at, updated_at) 
             VALUES (?, ?, ?, ?, ?, ?)`,
			payment.ToUserID, payment.FromUserID, payment.Amount, 0, createDate, createDate)

		if err != nil {
			log.Println("Error creating debt:", err)
			return response, err
		}
	}

	// Aktualisiere die Benutzerinformationen für beide beteiligten Benutzer
	err = d.UpdateUserDebtsAndCredits(payment.FromUserID)
	if err != nil {
		log.Println("Error updating sender's debt information:", err)
		return response, err
	}

	err = d.UpdateUserDebtsAndCredits(payment.ToUserID)
	if err != nil {
		log.Println("Error updating receiver's debt information:", err)
		return response, err
	}

	response.Success = true
	return response, nil
}
