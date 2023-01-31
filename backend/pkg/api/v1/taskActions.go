package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models"
	finLog2 "github.com/qrds1/Grats.Card/backend/pkg/common/models/finLog"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models/taskList"
	"github.com/qrds1/Grats.Card/backend/pkg/common/servicesAPI/AdMediaCards"
	"github.com/qrds1/Grats.Card/backend/pkg/common/tgBot"
	"net/http"
	"strconv"
	"time"
)

type CreateTaskCardCreationRequestBody struct {
	CardType string  `json:"cardType"`
	BIN      string  `json:"BIN"`
	Note     string  `json:"note"`
	Limit    float64 `json:"limit"`
}

func (h handler) CreateTaskCardCreation(c *gin.Context) {
	body := CreateTaskCardCreationRequestBody{}
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
	var neededAmount float64
	var cardCreateTax float64
	var count = 1
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
	neededAmount = float64(count) * (body.Limit + cardCreateTax)
	// check if user has enough money on the account
	if user.USDTBalance < neededAmount {
		printErrorMessage(c, "notEnoughBalance")
		return
	}
	initial := user.USDTBalance
	user.USDTBalance = user.USDTBalance - neededAmount
	// take off the card replenishment money

	h.DB.Save(&user)
	var task taskList.Task
	task.Status = 1 // ожидание
	task.Type = "CardCreation"
	task.UserID = user.ID
	task.Price = neededAmount
	task.InitialBalance = initial
	jsonData, _ := json.Marshal(&body)
	task.JsonData = string(jsonData)
	h.DB.Create(&task)
	tgBot.SendRequest(user)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

type CreateTaskCardReplenishmentRequestBody struct {
	CardId uint    `json:"cardId"`
	Amount float64 `json:"amount"`
}

func (h handler) CreateTaskCardReplenishment(c *gin.Context) {
	body := CreateTaskCardReplenishmentRequestBody{}
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
	// check if user has enough money on the account
	if user.USDTBalance < body.Amount {
		printErrorMessage(c, "notEnoughBalance")
		return
	}
	initial := user.USDTBalance
	user.USDTBalance = user.USDTBalance - body.Amount
	// take off the card replenishment money

	h.DB.Save(&user)
	var task taskList.Task
	task.Status = 1 // ожидание
	task.Type = "CardReplenishment"
	task.UserID = user.ID
	task.CardID = body.CardId
	task.Price = body.Amount
	task.InitialBalance = initial
	jsonData, _ := json.Marshal(&body)
	task.JsonData = string(jsonData)
	h.DB.Create(&task)
	tgBot.SendRequest(user)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

type TaskIdRequestBody struct {
	TaskId uint `json:"taskId"`
}

func (h handler) RepeatTask(c *gin.Context) {
	body := TaskIdRequestBody{}
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
	var taskPrevious taskList.Task
	h.DB.Find(&taskPrevious, body.TaskId)
	// check if user has enough money on the account
	if user.USDTBalance < taskPrevious.Price {
		printErrorMessage(c, "notEnoughBalance")
		return
	}
	initial := user.USDTBalance
	user.USDTBalance = user.USDTBalance - taskPrevious.Price
	// take off the card replenishment money
	h.DB.Save(&user)
	var task taskList.Task
	task.Status = 1 // ожидание
	task.Type = taskPrevious.Type
	task.UserID = user.ID
	task.CardID = taskPrevious.CardID
	task.Price = taskPrevious.Price
	task.InitialBalance = initial
	task.JsonData = taskPrevious.JsonData
	h.DB.Create(&task)
	tgBot.SendRequest(user)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (h handler) AdminGetTasks(c *gin.Context) {
	var tasks []taskList.Task
	if result := h.DB.Find(&tasks); result.Error != nil {
		printErrorMessage(c, "databaseError")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"tasks":   &tasks,
	})
}

func (h handler) GetTasks(c *gin.Context) {
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
	var tasks []taskList.Task
	if result := h.DB.Where("user_id = ?", user.ID).Find(&tasks); result.Error != nil {
		printErrorMessage(c, "databaseError")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"tasks":   &tasks,
	})
}

type CardCreationByTaskRequestBody struct {
	TaskId       int     `json:"taskId"`
	UserID       int     `json:"userID"`
	CardProvider string  `json:"cardProvider"`
	CardType     string  `json:"cardType"`
	Balance      string  `json:"balance"`
	BIN          string  `json:"BIN"`
	Note         string  `json:"note"`
	Limit        float64 `json:"limit"`
	Count        int     `json:"count"`
}

func (h handler) CardCreationByTask(c *gin.Context) {
	body := CardCreationByTaskRequestBody{}
	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}
	// get user info
	var user models.User
	if result := h.DB.First(&user, body.UserID); result.Error != nil {
		printErrorMessage(c, "userNotFound")
		return
	}
	var cardCreateTax float64
	var commission float64
	cardCreateTax = user.UniversalCost
	commission = user.UniversalCommission / 100
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
				commission = user.AdvertisingCommission / 100
				break
			}
		}
	}
	var task taskList.Task
	h.DB.Find(&task, task.ID)
	if task.Status != 1 {
		printErrorMessage(c, "statusError")
		return
	}
	//create all cards in cycle
	for i := 1; i <= body.Count; i++ {
		bin, _ := strconv.Atoi(body.BIN)
		var address models.Address
		if result := h.DB.Order("random()").Where("status = ?", 0).First(&address); result.Error != nil {
			printErrorMessage(c, "noAvailableAddresses")
			return
		}
		// generate useless data for card creation
		payload, addressString, name := generateRandomUserData(bin, body.Limit*(1-commission), address)
		// create a card
		response, err := AdMediaCards.CardCreate(payload)
		if err != nil {
			return
		}
		// log
		r, _ := json.Marshal(&response)
		fmt.Println(string(r))
		// log ended
		// check if card creation is approved by provider
		if !response.Result.IsApproved {
			printErrorMessage(c, "card is Not Approved")
			return
		}
		// disable used address
		address.Status = 1
		h.DB.Save(&address)
		// create card info in our system
		var card models.Card
		card.CardProvider = body.CardProvider
		card.ProviderId = response.Result.CardID
		card.CardNumber = "0000000000000000"
		card.CVV = "000"
		card.Balance = body.Limit * (1 - commission)
		currentTime := time.Now()
		card.IssuedDate = currentTime.Format("02/01/2006")
		card.Limit = body.Limit * (1 - commission)
		card.Spend = 0
		card.Status = 2
		card.Note = body.Note
		card.Address = addressString
		card.ExpirationDate = "00/00"
		card.Name = name
		card.UserId = user.ID
		h.DB.Create(&card)
		task.CardID = card.ID
		// create a transaction from card creation
		email1, _ := c.Get("email")
		email := email1.(string)
		var user2 models.User
		h.DB.Where("email = ?", email).First(&user2)
		initial := task.InitialBalance
		var transaction models.Transaction
		transaction.Status = 1
		transaction.Settled = true
		transaction.CardId = card.ID
		transaction.Provider = "Grats.Cards"
		transaction.UserId = user.ID
		transaction.Amount = cardCreateTax
		transaction.Commission = 0
		transaction.Transaction = ""
		transaction.Country = ""
		transaction.Merchant = ""
		transaction.Notes = "Card Creation"
		transaction.Initial = initial
		transaction.Total = initial - cardCreateTax
		transaction.Increase = false
		transaction.Balance = body.Balance
		transaction.ManagerId = user2.ID
		h.DB.Create(&transaction)
		// create a transaction from card replenishment
		initial = initial - cardCreateTax
		var transaction2 models.Transaction
		transaction2.Status = 1
		transaction2.Settled = true
		transaction2.CardId = card.ID
		transaction2.Provider = "Grats.Cards"
		transaction2.UserId = user.ID
		transaction2.Amount = body.Limit
		transaction2.Commission = commission * 100
		transaction2.Transaction = ""
		transaction2.Country = ""
		transaction2.Merchant = ""
		transaction2.Notes = "Card Replenishment"
		transaction2.Initial = initial
		transaction2.Total = initial - body.Limit
		transaction2.Increase = false
		transaction2.Balance = body.Balance
		transaction2.ManagerId = user2.ID
		h.DB.Create(&transaction2)
		var finLog finLog2.FinLog
		finLog.CreatorID = user2.ID
		finLog.UserID = user.ID
		finLog.Type = "CardCreation"
		finLog.TypeID = 0
		h.DB.Create(&finLog)
		var actionLog finLog2.CardCreation
		actionLog.CreatorID = user2.ID
		actionLog.UserID = user.ID
		actionLog.Type = "CardCreation"
		actionLog.MainID = finLog.ID
		actionLog.CardID = card.ID
		actionLog.TransactionID = transaction.ID
		h.DB.Create(&actionLog)
		finLog.TypeID = actionLog.ID
		h.DB.Save(&finLog)
	}
	task.Status = 3
	h.DB.Save(&task)
	fmt.Println("success", &user)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user":    &user,
	})
}

func (h handler) CardReplenishmentByTask(c *gin.Context) {
	body := TopUpCardBalanceRequestBody{}
	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}
	var task taskList.Task
	h.DB.Find(&task, task.ID)
	var card models.Card
	if result := h.DB.First(&card, body.CardID); result.Error != nil {
		printErrorMessage(c, "cardNotFound")
		return
	}
	var user models.User
	if result := h.DB.First(&user, card.UserId); result.Error != nil {
		printErrorMessage(c, "userNotFound")
		return
	}
	var commission float64
	commission = user.UniversalCommission / 100
	var suppliers []models.Supplier
	h.DB.Find(&suppliers)
	binFromCard := card.CardNumber[:6]
	for _, supplier := range suppliers {
		var bins []Bin
		err := json.Unmarshal([]byte(supplier.Bins), &bins)
		if err != nil {
			printErrorMessage(c, "bins error")
			return
		}
		for _, bin := range bins {
			if strconv.Itoa(bin.Bin) == binFromCard && !bin.IsUniversal {
				commission = user.AdvertisingCommission / 100
				break
			}
		}
	}
	if task.Status != 1 {
		printErrorMessage(c, "statusError")
		return
	}
	var payload AdMediaCards.ChangeCardLimitPayload
	payload.Type = "Limit"
	payload.Limit = card.Limit + body.Amount*(1-commission)
	response, err := AdMediaCards.ChangeCardLimit(card.ProviderId, payload)
	if err != nil {
		return
	}
	r, _ := json.Marshal(&response)
	fmt.Println(string(r))
	if !response.Result.IsApproved {
		printErrorMessage(c, "change is Not Approved")
		return
	}
	email1, _ := c.Get("email")
	email := email1.(string)
	var user2 models.User
	h.DB.Where("email = ?", email).First(&user2)
	initial := task.InitialBalance
	card.Limit = card.Limit + body.Amount*(1-commission)
	card.Balance = card.Balance + body.Amount*(1-commission)
	h.DB.Save(&card)
	var transaction models.Transaction
	transaction.Status = 1
	transaction.Settled = true
	transaction.CardId = card.ID
	transaction.Provider = "Grats.Cards"
	transaction.UserId = user.ID
	transaction.Amount = body.Amount
	transaction.Commission = commission * 100
	transaction.Transaction = ""
	transaction.Country = ""
	transaction.Merchant = ""
	transaction.Notes = "Card Replenishment"
	transaction.Initial = initial
	transaction.Total = initial - body.Amount
	transaction.Increase = false
	transaction.Balance = body.Balance
	transaction.ManagerId = user2.ID
	h.DB.Create(&transaction)
	task.Status = 3
	h.DB.Save(&task)
	fmt.Println("success", &card)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}

func (h handler) CancelTask(c *gin.Context) {
	body := TaskIdRequestBody{}
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
	var task taskList.Task
	h.DB.Find(&task, body.TaskId)
	// check if user has enough money on the account
	if task.Status != 1 {
		printErrorMessage(c, "TaskIsNotInProgress")
		return
	}
	user.USDTBalance = user.USDTBalance + task.Price
	h.DB.Save(&user)
	task.Status = 2 // отмена
	h.DB.Save(&task)
	tgBot.SendRequest(user)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
