package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name                          string    `gorm:"not null" json:"name"`                 // Полное имя
	Email                         string    `gorm:"unique;not null" json:"email"`         // Электронная почта
	PasswordHash                  string    `gorm:"not null" json:"passwordHash"`         // Пароль
	Country                       string    `gorm:"not null" json:"country"`              // Страна проживания
	Messenger                     string    `gorm:"unique;not null" json:"messenger"`     // Мессенджер для связи
	Role                          int       `gorm:"not null;default:1" json:"role"`       // Роль пользователя
	Status                        int       `gorm:"not null;default:0" json:"status"`     // Статус активации аккаунта
	Currency                      string    `gorm:"not null;default:USD" json:"currency"` // Основная валюта
	USDTBalance                   float64   `gorm:"not null;default:0.00" json:"USDTBalance"`
	BTCBalance                    float64   `gorm:"not null;default:0.00" json:"BTCBalance"`
	USDBalance                    float64   `gorm:"not null;default:0.00" json:"USDBalance"`
	EURBalance                    float64   `gorm:"not null;default:0.00" json:"EURBalance"`
	LastIP                        string    `json:"lastIP"`
	LastActionDate                time.Time `json:"lastActionDate"`
	PopUpCommission               float64   `gorm:"not null;default:1" json:"popUpCommission"`
	UniversalCommission           float64   `gorm:"not null;default:8" json:"universalCommission"`
	UniversalCost                 float64   `gorm:"not null;default:15" json:"universalCost"`
	AdvertisingCommission         float64   `gorm:"not null;default:6" json:"advertisingCommission"`
	AdvertisingCost               float64   `gorm:"not null;default:10" json:"advertisingCost"`
	TimeZone                      string    `gorm:"not null;default:'(UTC) Great Britain'" json:"timeZone"`
	UniversalMonthlyMaintenance   float64   `gorm:"not null;default:15" json:"universalMonthlyMaintenance"`
	AdvertisingMonthlyMaintenance float64   `gorm:"not null;default:10" json:"advertisingMonthlyMaintenance"`
}
