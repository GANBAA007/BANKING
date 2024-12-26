package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	ClientID      uint
	AccountNo     string
	Balance       float64
	ClientSurname string
	ClientName    string
}
