package models

import (
	"time"

	"gorm.io/gorm"
)

type Savings struct {
	gorm.Model     //
	AccountNo      string
	ClientID       uint      //
	MonthlyPayment float64   //
	PaymentDate    time.Time //
	MinAmount      float64   //
	Balance        float64   //
	ExpDate        time.Time //
	InterestRate   float64   //
	InterestProfit float64
}
