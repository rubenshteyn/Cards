package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	Status          int       `gorm:"not null" json:"status"` // 0 - disabled, 1 - active
	Settled         bool      `gorm:"not null" json:"settled"`
	Provider        string    `json:"provider"`
	ProviderId      uint      `json:"providerId"`
	CardId          uint      `gorm:"not null" json:"cardId"`
	UserId          uint      `gorm:"not null" json:"userId"`
	Amount          float64   `gorm:"not null" json:"amount"`
	Country         string    `json:"country"`
	Merchant        string    `json:"merchant"`
	ManagerId       uint      `json:"managerId"`
	Increase        bool      `json:"increase"`
	Commission      float64   `json:"commission"`
	Initial         float64   `json:"initial"`
	Total           float64   `json:"total"`
	TransactionDate time.Time `json:"TransactionDate"`
	Transaction     string    `json:"transaction"`
	Notes           string    `json:"notes"`
	Balance         string    `json:"balance"`
}
