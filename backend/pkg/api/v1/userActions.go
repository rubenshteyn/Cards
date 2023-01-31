package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qrds1/Grats.Card/backend/pkg/common/hashPassword"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models"
	"net/http"
)

func (h handler) GetUserInfo(c *gin.Context) {
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
	var cards []models.Card
	if result := h.DB.Where("user_id = ?", user.ID).Find(&cards); result.Error != nil {
		printErrorMessage(c, "databaseError")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user":    &user,
		"cards":   &cards,
	})
}

type ChangeUserRequestBody struct {
	Type        string `json:"type"`
	OldPassword string `json:"oldPassword"`
	Password    string `json:"password"`
}

func (h handler) ChangeUserInfo(c *gin.Context) {
	body := ChangeUserRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}

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
	if body.Type == "Password" {
		check := hashPassword.CheckPasswordHash(body.OldPassword, user.PasswordHash)
		if !check {
			printErrorMessage(c, "incorrectPassword")
			return
		}
		user.PasswordHash, _ = hashPassword.HashPassword(body.Password)
	}
	if result := h.DB.Save(&user); result.Error != nil {
		printErrorMessage(c, "userNotUpdated")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
