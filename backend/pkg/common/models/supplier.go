package models

import (
	"gorm.io/gorm"
)

type Supplier struct {
	gorm.Model
	Status int    `gorm:"not null" json:"status"` // 0 - disabled, 1 - active
	Name   string `gorm:"not null" json:"name"`
	Bins   string `gorm:"not null" json:"bins"`
}
