package models

import (
	"gorm.io/gorm"
	"time"
)

type Card struct {
	gorm.Model
	Status         int       `json:"status"` // 0 - disabled, 1 - active
	CardProvider   string    `json:"cardProvider"`
	ProviderId     int       `json:"providerId"`
	CardNumber     string    `json:"cardNumber"`
	ExpirationDate string    `json:"expirationDate"`
	IssuedDate     string    `json:"issuedDate"`
	DeactivateDate time.Time `json:"deactivateDate"`
	CVV            string    `json:"CVV"`
	Note           string    `json:"note"`
	Limit          float64   `gorm:"default:0" json:"limit"`
	Balance        float64   `gorm:"default:0" json:"balance"`
	Spend          float64   `gorm:"default:0" json:"spend"`
	Name           string    `json:"name"`
	Address        string    `json:"address"`
	UserId         uint      `json:"userId"`
}
