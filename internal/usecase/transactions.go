package usecase

import (
	"bytes"
	"context"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/andrew221293/transactions-stori-challenge/internal/entity"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s StoriUseCase) ValidateTransaction(
	ctx context.Context,
	transactions []entity.Transaction) (entity.TransactionHistory, error) {
	var debit []float64
	var credit []float64
	var total []float64
	userID := transactions[0].UserID
	mapPerMonth := make(map[string]int)

	for _, v := range transactions {
		m := getMonth(v.Date)
		month := monthsName(m)
		totalPerMonth := mapPerMonth[month]
		mapPerMonth[month] = totalPerMonth + 1

		typeTransaction := v.Transaction[0:1]
		if typeTransaction == "-" {
			deb := strings.TrimLeft(v.Transaction, "-")
			d, err := strconv.ParseFloat(deb, 64)
			if err != nil {
				return entity.TransactionHistory{}, entity.CustomError{
					Err:      err,
					HTTPCode: http.StatusBadRequest,
					Code:     "d3e601d5-6482-49d6-9996-3d94cbaf740a",
				}
			}
			debit = append(debit, d)
		} else {
			cred := strings.TrimLeft(v.Transaction, "+")
			c, err := strconv.ParseFloat(cred, 64)
			if err != nil {
				return entity.TransactionHistory{}, entity.CustomError{
					Err:      err,
					HTTPCode: http.StatusBadRequest,
					Code:     "f31d63ed-34fa-453e-8b2f-e0e13cafe5a0",
				}
			}
			credit = append(credit, c)
		}

		totalTransactions := strings.TrimLeft(v.Transaction, "+")
		t, err := strconv.ParseFloat(totalTransactions, 64)
		if err != nil {
			return entity.TransactionHistory{}, entity.CustomError{
				Err:      err,
				HTTPCode: http.StatusBadRequest,
				Code:     "f31d63ed-34fa-453e-8b2f-e0e13cafe5a0",
			}
		}
		total = append(total, t)
	}

	totalBalance := totalCharges(total)
	totalDebitCharges := totalCharges(debit)
	totalCreditCharges := totalCharges(credit)

	user, err := s.Store.GetOneUser(ctx, userID)
	if err != nil {
		return entity.TransactionHistory{}, err
	}

	transaction, err := sendEmail(mapPerMonth, totalDebitCharges, totalCreditCharges, totalBalance, user)
	if err != nil {
		return entity.TransactionHistory{}, err
	}

	err = s.Store.InserTransactionHistory(ctx, transaction)
	if err != nil {
		return entity.TransactionHistory{}, err
	}

	return transaction, nil
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

func sendEmail(
	m map[string]int,
	totalDebit,
	totalCredit,
	total float64,
	user entity.User) (entity.TransactionHistory, error) {
	var transaction entity.TransactionHistory
	var perMonths []entity.TransactionsPerMonth
	for k, v := range m {
		perMonth := entity.TransactionsPerMonth{}
		perMonth.Month = k
		perMonth.Total = v
		perMonths = append(perMonths, perMonth)
	}

	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	t.Execute(&body, struct {
		Balance      float64
		Transactions []entity.TransactionsPerMonth
		Debit        float64
		Credit       float64
	}{
		Balance:      total,
		Transactions: perMonths,
		Debit:        totalDebit,
		Credit:       totalCredit,
	})

	from := mail.NewEmail("Andres Quintero", "storiandresromo@gmail.com")
	subject := "Transacciones"
	to := mail.NewEmail("Andres Romo", user.Email)
	plainTextContent := "story challenge"
	htmlContent := body.String()
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	if err != nil {
		return transaction, entity.CustomError{
			Err:      err,
			HTTPCode: http.StatusBadRequest,
			Code:     "2fe11db8-e7fe-445b-8309-24b8b3ea2ecf",
		}
	}

	transaction.ID = primitive.NewObjectID().Hex()
	transaction.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	transaction.Transactions = perMonths
	transaction.Balance = total
	transaction.Debit = totalDebit
	transaction.Credit = totalCredit

	return transaction, nil
}
