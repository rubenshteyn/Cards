package finLog

import (
	"gorm.io/gorm"
)

type FinLog struct {
	gorm.Model
	CreatorID uint   `json:"creatorID"` // user who did the action
	UserID    uint   `json:"userID"`    // user which was affected by the action
	Type      string `json:"type"`      // action type "cardCreation CardChange CardReplenishment AccountReplenishment etc"
	TypeID    uint   `json:"typeID"`    // action table ID
}
