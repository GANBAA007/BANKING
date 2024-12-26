package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Surname  string `gorm:"not null"`
	RegNo    string `gorm:"not null"`
	PhoneNo  string `gorm:"not null"`
	Password string `gorm:"not null"`
}
