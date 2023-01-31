package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goombaio/namegenerator"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models"
	finLog2 "github.com/qrds1/Grats.Card/backend/pkg/common/models/finLog"
	"github.com/qrds1/Grats.Card/backend/pkg/common/servicesAPI/AdMediaCards"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (h handler) AdminGetCards(c *gin.Context) {
	var cards []models.Card
	if result := h.DB.Find(&cards); result.Error != nil {
		printErrorMessage(c, "databaseError")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"cards":   &cards,
	})
}

type CreateCardsRequestBody struct {
	UserID       int     `json:"userID"`
	CardProvider string  `json:"cardProvider"`
	CardType     string  `json:"cardType"`
	Balance      string  `json:"balance"`
	BIN          string  `json:"BIN"`
	Note         string  `json:"note"`
	Limit        float64 `json:"limit"`
	Count        int     `json:"count"`
}

func (h handler) CreateCards(c *gin.Context) {
	body := CreateCardsRequestBody{}
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
	// get needed amount of money to create and replenish all the cards
	var neededAmount float64
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
		// take off the card creation money
		initial := user.USDTBalance
		user.USDTBalance = user.USDTBalance - cardCreateTax
		h.DB.Save(&user)
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
		// create a transaction from card creation
		email1, _ := c.Get("email")
		email := email1.(string)
		var user2 models.User
		h.DB.Where("email = ?", email).First(&user2)
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
		transaction.Total = user.USDTBalance
		transaction.Increase = false
		transaction.Balance = body.Balance
		transaction.ManagerId = user2.ID
		h.DB.Create(&transaction)
		// take off the card replenishment money
		initial = user.USDTBalance
		user.USDTBalance = user.USDTBalance - body.Limit
		h.DB.Save(&user)
		// create a transaction from card replenishment
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
		transaction2.Total = user.USDTBalance
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
	fmt.Println("success", &user)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user":    &user,
	})
}

func generateRandomUserData(bin int, limit float64, address models.Address) (AdMediaCards.CardCreatePayload, string, string) {
	var payload AdMediaCards.CardCreatePayload
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	name := nameGenerator.Generate()
	nameSlice := strings.Split(name, "-")
	caser := cases.Title(language.AmericanEnglish)
	firstName := caser.String(nameSlice[0])
	lastName := caser.String(nameSlice[1])
	name = firstName + " " + lastName
	payload.BIN = bin
	payload.Limit = limit
	payload.CountryIso = "US"
	payload.FirstName = firstName
	payload.LastName = lastName
	payload.Address1 = address.Address1
	payload.City = address.City
	payload.State = address.State
	payload.Zip = address.Zip
	rand.Seed(time.Now().UnixNano())
	months := rand.Intn(26) + 9
	now := time.Now()
	after := now.AddDate(0, months, 0)
	expYear := after.Year()
	month := after.Month()
	expMonth := int(month)
	payload.ExpMonth = strconv.Itoa(expMonth)
	if expMonth < 10 {
		payload.ExpMonth = "0" + strconv.Itoa(expMonth)
	}
	payload.ExpYear = strconv.Itoa(expYear)
	payload.PhoneNumber = address.PhoneNumber
	addressString := address.Address1 + ", " + address.City + ", " + address.State + " " + address.Zip + ", us"
	return payload, addressString, name
}

type TopUpCardBalanceRequestBody struct {
	CardID  int     `json:"cardID"`
	Balance string  `json:"balance"`
	Amount  float64 `json:"amount"`
	// CardIDs []int  `json:"cardIDs"`
}

type Bin struct {
	Bin         int    `json:"bin"`
	Active      bool   `json:"active"`
	IsUniversal bool   `json:"isUniversal"`
	Supplier    string `json:"supplier"`
}

func (h handler) TopUpCardBalance(c *gin.Context) {
	body := TopUpCardBalanceRequestBody{}
	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}
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
	var neededAmount float64
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
	neededAmount = body.Amount
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
	initial := user.USDTBalance
	user.USDTBalance = user.USDTBalance - body.Amount
	h.DB.Save(&user)
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
	transaction.Total = user.USDTBalance
	transaction.Increase = false
	transaction.Balance = body.Balance
	transaction.ManagerId = user2.ID
	h.DB.Create(&transaction)
	fmt.Println("success", &card)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}

type ChangeCardStatusRequestBody struct {
	CardID int  `json:"cardID"`
	Status bool `json:"status"`
}

func (h handler) ChangeCardStatus(c *gin.Context) {
	body := ChangeCardStatusRequestBody{}
	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}
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
	var payload AdMediaCards.ChangeCardStatusPayload
	payload.Type = "Status"
	payload.Status = body.Status
	response, err := AdMediaCards.ChangeCardStatus(card.ProviderId, payload)
	if err != nil {
		return
	}
	r, _ := json.Marshal(&response)
	fmt.Println(string(r))
	if !response.Result.IsApproved {
		printErrorMessage(c, "change is Not Approved")
		return
	}
	if body.Status {
		card.Status = 1
		card.DeactivateDate = time.Unix(0, 0)
	} else {
		card.Status = 0
		card.DeactivateDate = time.Now()
	}
	h.DB.Save(&card)
	fmt.Println("success", &card)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}

type ChangeCardInfoRequestBody struct {
	CardID    int    `json:"cardID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Address1  string `json:"Address1"`
	City      string `json:"City"`
	State     string `json:"State"`
	Country   string `json:"Country"`
	Zip       string `json:"Zip"`
}

func (h handler) ChangeCardInfo(c *gin.Context) {
	body := ChangeCardInfoRequestBody{}
	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}
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
	var payload AdMediaCards.ChangeCardNameAddressPayload
	payload.Type = "NameAddress"
	payload.FirstName = body.FirstName
	payload.LastName = body.LastName
	payload.Address1 = body.Address1
	payload.City = body.City
	payload.State = body.State
	payload.Country = body.Country
	payload.Zip = body.Zip
	response, err := AdMediaCards.ChangeCardNameAddress(card.ProviderId, payload)
	if err != nil {
		return
	}
	r, _ := json.Marshal(&response)
	fmt.Println(string(r))
	if !response.Result.IsApproved {
		printErrorMessage(c, "change is Not Approved")
		return
	}
	fmt.Println("success", &card)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}

type ChangeCardLimitRequestBody struct {
	CardID int     `json:"cardID"`
	Limit  float64 `json:"limit"`
}

func (h handler) ChangeCardLimit(c *gin.Context) {
	body := ChangeCardLimitRequestBody{}
	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}
	var card models.Card
	if result := h.DB.First(&card, body.CardID); result.Error != nil {
		printErrorMessage(c, "cardNotFound")
		return
	}
	var payload AdMediaCards.ChangeCardLimitPayload
	payload.Type = "Limit"
	payload.Limit = body.Limit
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
	card.Limit = body.Limit
	h.DB.Save(&card)
	fmt.Println("success", &card)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}
