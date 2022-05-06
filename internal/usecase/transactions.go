package usecase

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/andrew221293/transactions-stori-challenge/internal/entity"
)

func (s StoriUseCase) ValidateTransaction(transactions []entity.Transaction) error {
	var debit []float64
	var credit []float64

	months := make(map[string]string)

	for _, v := range transactions {
		m := getMonth(v.Date)
		month := monthsName(m)
		months[m] = month
		typeTransaction := v.Transaction[0:1]
		if typeTransaction == "-" {
			deb := strings.TrimLeft(v.Transaction, "-")
			d, err := strconv.ParseFloat(deb, 64)
			if err != nil {
				return entity.CustomError{
					Err:      err,
					HTTPCode: http.StatusBadRequest,
					Code:     "d3e601d5-6482-49d6-9996-3d94cbaf740a",
				}
			}
			debit = append(debit, d)
		}
		cred := strings.TrimLeft(v.Transaction, "+")
		c, err := strconv.ParseFloat(cred, 64)
		if err != nil {
			return entity.CustomError{
				Err:      err,
				HTTPCode: http.StatusBadRequest,
				Code:     "f31d63ed-34fa-453e-8b2f-e0e13cafe5a0",
			}
		}
		credit = append(credit, c)
	}

	totalDebitCharges := totalCharges(debit)
	totalCreditCharges := totalCharges(credit)

	err := sendEmail(months, totalDebitCharges, totalCreditCharges)
	if err != nil {
		return err
	}

	return nil
}

func getMonth(date string) string {
	month := date[0:1]
	return month
}

func monthsName(m string) string {
	var month string
	switch m {
	case "1":
		month = "January"
	case "2":
		month = "February"
	case "3":
		month = "March"
	case "4":
		month = "April"
	case "5":
		month = "May"
	case "6":
		month = "June"
	case "7":
		month = "July"
	case "8":
		month = "August"
	case "9":
		month = "September"
	case "10":
		month = "October"
	case "11":
		month = "November"
	case "12":
		month = "December"
	default:
		month = "not valid"
	}

	return month
}

func totalCharges(transactions []float64) float64 {
	var total float64

	for _, v := range transactions {
		total = total + v
	}

	return total
}

func sendEmail(m map[string]string, totalDebit, totalCredit float64) error {
	// Sender data.
	from := "storiAndresRomo@gmail.com"
	password := "4lq43d45INC"

	// Receiver email address.
	to := []string{
		"andres_romo93@hotmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Balance      string
		Transactions []entity.TransactionsPerMonth
		Debit        string
		Credit       string
	}{
		//TODO Make HTML
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return entity.CustomError{
			Err:      err,
			HTTPCode: http.StatusInternalServerError,
			Code:     "2ee64a34-4cd4-49a3-8752-6b454de22a3a",
		}
	}
	fmt.Println("Email Sent!")

	return nil
}
