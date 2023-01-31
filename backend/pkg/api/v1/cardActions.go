package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models"
	"github.com/qrds1/Grats.Card/backend/pkg/common/tgBot"
	"net/http"
	"strconv"
)

func (h handler) GetCards(c *gin.Context) {
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
		"cards":   &cards,
	})
}

type UserCreateCardRequestBody struct {
	CardType string  `json:"cardType"`
	Balance  string  `json:"balance"`
	BIN      string  `json:"BIN"`
	Note     string  `json:"note"`
	Action   string  `json:"action"`
	Limit    float64 `json:"limit"`
	Count    int     `json:"count"`
}

func (h handler) UserCreateCard(c *gin.Context) {
	body := UserCreateCardRequestBody{}
	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}
	// get user info
	email1, exists := c.Get("email")
	if !exists {
		printErrorMessage(c, "Cannot check email")
		return
	}
	body.Action = "создание карты"
	email := email1.(string)
	var user models.User
	if result := h.DB.Where("email = ?", email).First(&user); result.Error != nil {
		printErrorMessage(c, "Unauthorized, user not exists")
		return
	}
	// get needed amount of money to create and replenish all the cards
	var neededAmount float64
	var cardCreateTax float64
	cardCreateTax = user.UniversalCost
	var suppliers []models.Supplier
	h.DB.Find(&suppliers)
	binFromCard := body.BIN
	for _, supplier := range suppliers {
		var bins []Bin
		err := json.Unmarshal([]byte(supplier.Bins), &bins)
		if err != nil {
			printErrorMessage(c, "bins error")
			return
		}
		for _, bin := range bins {
			if strconv.Itoa(bin.Bin) == binFromCard && !bin.IsUniversal {
				cardCreateTax = user.AdvertisingCost
				break
			}
		}
	}
	neededAmount = float64(body.Count) * (body.Limit + cardCreateTax)
	// check if user has enough money on the account
	if body.Balance == "USDT" {
		if user.USDTBalance < neededAmount {
			printErrorMessage(c, "notEnoughBalance")
			return
		}
	} else {
		// todo: other balances
		printErrorMessage(c, "paying not from USDT balance is disabled for now")
		return
	}
	tgBot.SendRequest(user)
	fmt.Println("success", &user)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user":    &user,
	})
}
