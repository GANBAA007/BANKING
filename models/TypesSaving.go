package models

import (
	"time"

	"gorm.io/gorm"
)

type Savings struct {
	gorm.Model
	ClientID       uint
	MonthlyPayment float64
	MinAmount      float64
	IsChecking     bool
	IsYieldofYield bool
	Balance        float64
	ExpDate        time.Time
}
