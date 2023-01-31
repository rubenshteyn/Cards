package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Status      int    `gorm:"not null;default:0" json:"status"` // 0 - not used, 1 - used
	CountryIso  string `gorm:"not null;default:US" json:"countryIso"`
	Address1    string `gorm:"not null" json:"address1"`
	City        string `gorm:"not null" json:"city"`
	State       string `gorm:"not null" json:"state"`
	Zip         string `gorm:"not null" json:"zip"`
	PhoneNumber string `gorm:"not null" json:"phoneNumber"`
}
