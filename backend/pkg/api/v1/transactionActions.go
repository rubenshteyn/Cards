package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models"
	"net/http"
)

type CardIDBody struct {
	CardID uint `json:"cardID"`
}

func (h handler) GetTransactions(c *gin.Context) {
	body := CardIDBody{}
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
	var transactions []models.Transaction
	if body.CardID == 0 {
		if result := h.DB.Where("user_id = ?", user.ID).Find(&transactions); result.Error != nil {
			printErrorMessage(c, "databaseError")
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success":      true,
			"transactions": &transactions,
		})
		return
	}
	if user.Role == 2 {
		if result := h.DB.Where("card_id = ?", body.CardID).Find(&transactions); result.Error != nil {
			printErrorMessage(c, "databaseError")
			return
		}
	} else {
		if result := h.DB.Where("card_id = ? && user_id = ?", body.CardID, user.ID).Find(&transactions); result.Error != nil {
			printErrorMessage(c, "databaseError")
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"transactions": &transactions,
	})
}

func (h handler) GetAllTransactions(c *gin.Context) {
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
	var transactions []models.Transaction
	if result := h.DB.Where("user_id = ?", user.ID).Find(&transactions); result.Error != nil {
		printErrorMessage(c, "databaseError")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"transactions": &transactions,
	})
}
