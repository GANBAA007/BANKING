package models

import "gorm.io/gorm"

type Wallet struct {
	AccountNo string
	// Client Id FK
	gorm.Model
}
