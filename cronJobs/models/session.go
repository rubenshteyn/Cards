package models

import (
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	AuthenticationToken string `gorm:"unique;not null" json:"authenticationToken"`
	Status              int    `gorm:"not null" json:"status"` // 0 - disabled, 1 - active, 2 - admin, 3 - expired, 4 - from admin
	UserId              uint   `gorm:"not null" json:"userId"`
	Location            string `gorm:"not null" json:"location"`
}
