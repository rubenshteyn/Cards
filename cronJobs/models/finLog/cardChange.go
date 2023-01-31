package finLog

import (
	"gorm.io/gorm"
)

type CardChange struct {
	gorm.Model
	CreatorID       uint   `json:"creatorID"` // user who did the action
	UserID          uint   `json:"userID"`    // user which was affected by the action
	Type            string `json:"type"`      // action type "cardCreation CardChange CardReplenishment AccountReplenishment etc"
	MainID          uint   `json:"mainID"`    // main table ID
	CardID          uint   `json:"cardID"`    // card ID
	ChangeType      string `json:"changeType"`
	WhatChangedJSON string `json:"whatChangedJSON"`
}
