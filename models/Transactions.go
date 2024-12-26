package models

import "gorm.io/gorm"

type Trans struct {
	gorm.Model
	Dest_acc      string
	Src_acc       string
	Amount        float64
	Approval_code uint
}
