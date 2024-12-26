package models

import "gorm.io/gorm"

type Savings struct {
	gorm.Model
	ClientID uint
}
