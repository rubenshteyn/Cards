package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models/finLog"
	"net/http"
)

func (h handler) AdminGetFinLogs(c *gin.Context) {
	var finLogs []finLog.FinLog
	if result := h.DB.Find(&finLogs); result.Error != nil {
		printErrorMessage(c, "databaseError")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"finLogs": &finLogs,
	})
}

type GetActionLogRequestBody struct {
	Type     string `json:"type"`
	ActionID uint   `json:"actionID"`
}

func (h handler) AdminGetActionLog(c *gin.Context) {
	body := GetActionLogRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}
	if body.Type == "CardCreation" {
		var cc finLog.CardCreation
		if result := h.DB.First(&cc, body.ActionID); result.Error != nil {
			printErrorMessage(c, "userNotExist")
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    &cc,
		})
		return
	}
	if body.Type == "Replenishment" {
		var cc finLog.Replenishment
		if result := h.DB.First(&cc, body.ActionID); result.Error != nil {
			printErrorMessage(c, "userNotExist")
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    &cc,
		})
		return
	}
	if body.Type == "CardChange" {
		var cc finLog.CardChange
		if result := h.DB.First(&cc, body.ActionID); result.Error != nil {
			printErrorMessage(c, "userNotExist")
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    &cc,
		})
		return
	}
	printErrorMessage(c, "Type is not supported yet")
}
