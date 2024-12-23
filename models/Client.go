package models

import "gorm.io/gorm"

type Client struct {
	Name    string
	Surname string
	gorm.Model
	RegNo   string
	PhoneNo string
}
