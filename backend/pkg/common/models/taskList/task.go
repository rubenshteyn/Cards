package taskList

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ManagerID      uint    `json:"creatorID"`               // user who did the task
	UserID         uint    `json:"userID"`                  // user who created the task
	CardID         uint    `gorm:"default:0" json:"cardID"` // card
	Type           string  `json:"type"`                    // action type "cardCreation CardChange CardReplenishment AccountReplenishment etc"
	JsonData       string  `json:"jsonData"`                // json data with all needed info about the action
	Price          float64 `json:"price"`                   // price of the action
	InitialBalance float64 `json:"initialBalance"`          // initialBalance
	Status         int     `json:"status"`                  // action status (0 - rejected/1 - waiting/2 - canceled/3 - completed)
}
