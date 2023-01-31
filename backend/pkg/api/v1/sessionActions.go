package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models"
	"net/http"
)

type SessionIDBody struct {
	SessionID int  `json:"sessionID"`
	DeleteAll bool `json:"deleteAll"`
}

func (h handler) DeleteSessions(c *gin.Context) {
	body := SessionIDBody{}
	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}

	if body.DeleteAll {
		email1, exists := c.Get("email")
		if !exists {
			printErrorMessage(c, "Cannot check email")
			return
		}
		email := email1.(string)
		var user models.User
		if result := h.DB.Where("email = ?", email).First(&user); result.Error != nil {
			printErrorMessage(c, "Unauthorized, user not exists")
			return
		}
		var session models.Session
		fmt.Println(user.ID, body.SessionID)
		if result := h.DB.Delete(&session, "user_id = ? AND id != ?", user.ID, body.SessionID); result.Error != nil {
			printErrorMessage(c, "databaseError")
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}
	var session models.Session
	if result := h.DB.Delete(&session, body.SessionID); result.Error != nil {
		printErrorMessage(c, "databaseError")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (h handler) GetSessions(c *gin.Context) {
	email1, exists := c.Get("email")
	if !exists {
		printErrorMessage(c, "Cannot check email")
		return
	}
	email := email1.(string)
	var user models.User
	if result := h.DB.Where("email = ?", email).First(&user); result.Error != nil {
		printErrorMessage(c, "userNotFound")
		return
	}
	var sessions []models.Session
	forbiddenStatus := 4
	if user.Role == 2 {
		forbiddenStatus = 5
	}
	if result := h.DB.Unscoped().Where("user_id = ? AND status != ?", user.ID, forbiddenStatus).Find(&sessions); result.Error != nil {
		printErrorMessage(c, "databaseError")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"sessions": &sessions,
	})
}
