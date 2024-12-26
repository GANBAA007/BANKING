package models

import "gorm.io/gorm"

type Loan struct {
	gorm.Model
	ClientID    uint    `gorm:"not null"`
	Amount      float64 `gorm:"not null"`
	Interest    float64 `gorm:"not null"`
	Is_compound bool    `gorm:"not null"`
	Is_mortgage bool    `gorm:"not null"`
}
