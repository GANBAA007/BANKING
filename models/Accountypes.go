package models

import "gorm.io/gorm"

type AccTypes struct {
	gorm.Model
	Type    string
	Percent float64
}
