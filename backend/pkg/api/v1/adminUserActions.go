package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qrds1/Grats.Card/backend/pkg/common/hashPassword"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models"
	"github.com/qrds1/Grats.Card/backend/pkg/common/servicesAPI/googleReCaptcha"
	"github.com/qrds1/Grats.Card/backend/pkg/common/token"
	"net/http"
	"time"
)

func (h handler) AdminGetUsers(c *gin.Context) {
	var users []models.User
	if result := h.DB.Find(&users); result.Error != nil {
		printErrorMessage(c, "databaseError")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"users":   &users,
	})
}

type AdminAddUserRequestBody struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Country   string `json:"country"`
	Messenger string `json:"messenger"`
	Status    int    `json:"status"`
	Role      int    `json:"role"`
}

func (h handler) AdminAddUser(c *gin.Context) {
	body := AdminAddUserRequestBody{}
	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}
	var user models.User
	user.Name = body.Name
	user.Email = body.Email
	user.PasswordHash, _ = hashPassword.HashPassword(body.Password)
	user.Country = body.Country
	user.Messenger = body.Messenger
	ip := c.ClientIP()
	user.LastIP = ip
	user.LastActionDate = time.Now()
	user.Status = body.Status
	user.Role = body.Role
	if result := h.DB.Create(&user); result.Error != nil {
		printErrorMessage(c, "emailAlreadyRegistered")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user":    &user,
	})
}

type AdminUpdateUserRequestBody struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Country   string `json:"country"`
	Messenger string `json:"messenger"`
	Status    int    `json:"status"`
	Role      int    `json:"role"`
	ID        uint   `json:"id"`
}

func (h handler) AdminUpdateUser(c *gin.Context) {
	body := AdminUpdateUserRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}

	var user models.User
	h.DB.First(&user, body.ID)
	if body.Name != user.Name {
		user.Name = body.Name
	}
	if body.Email != user.Email {
		user.Email = body.Email
	}
	if len(body.Password) > 0 {
		user.PasswordHash, _ = hashPassword.HashPassword(body.Password)
	}
	if body.Country != user.Country {
		user.Country = body.Country
	}
	if body.Messenger != user.Messenger {
		user.Messenger = body.Messenger
	}
	if body.Status != user.Status {
		user.Status = body.Status
		if user.Status != 2 {
			var session models.Session
			h.DB.Delete(&session, "user_id = ?", user.ID)
		}
	}
	if body.Role != user.Role {
		user.Role = body.Role
	}
	if result := h.DB.Save(&user); result.Error != nil {
		printErrorMessage(c, "userNotUpdated "+result.Error.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user":    &user,
	})
}

type AdminUpdateUsersCommissionRequestBody struct {
	UniversalCommission           float64 `json:"universalCommission"`
	UniversalCost                 float64 `json:"universalCost"`
	AdvertisingCommission         float64 `json:"advertisingCommission"`
	AdvertisingCost               float64 `json:"advertisingCost"`
	UniversalMonthlyMaintenance   float64 `json:"universalMonthlyMaintenance"`
	AdvertisingMonthlyMaintenance float64 `json:"advertisingMonthlyMaintenance"`
	ID                            uint    `json:"id"`
}

func (h handler) AdminUpdateUsersCommission(c *gin.Context) {
	body := AdminUpdateUsersCommissionRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}

	var user models.User
	h.DB.First(&user, body.ID)
	user.UniversalCommission = body.UniversalCommission
	user.AdvertisingCommission = body.AdvertisingCommission
	user.UniversalCost = body.UniversalCost
	user.AdvertisingCost = body.AdvertisingCost
	user.UniversalMonthlyMaintenance = body.UniversalMonthlyMaintenance
	user.AdvertisingMonthlyMaintenance = body.AdvertisingMonthlyMaintenance
	if result := h.DB.Save(&user); result.Error != nil {
		printErrorMessage(c, "userNotUpdated "+result.Error.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user":    &user,
	})
}

type AdminTopUpBalanceRequestBody struct {
	ID          uint    `json:"id"`
	Balance     string  `json:"balance"`
	Amount      float64 `json:"amount"`
	Sum         float64 `json:"sum"`
	Commission  float64 `json:"commission"`
	Transaction string  `json:"transaction"`
	Notes       string  `json:"notes"`
}

func (h handler) AdminTopUpBalance(c *gin.Context) {
	body := AdminTopUpBalanceRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}

	var user models.User
	h.DB.First(&user, body.ID)
	var initial float64
	var total float64
	switch body.Balance {
	case "USDT":
		initial = user.USDTBalance
		user.USDTBalance += body.Amount
		total = user.USDTBalance
	case "USD":
		initial = user.USDBalance
		user.USDBalance += body.Amount
		total = user.USDBalance
	case "EUR":
		initial = user.EURBalance
		user.EURBalance += body.Amount
		total = user.EURBalance
	case "BTC":
		initial = user.BTCBalance
		user.BTCBalance += body.Amount
		total = user.BTCBalance
	}
	if result := h.DB.Save(&user); result.Error != nil {
		printErrorMessage(c, "userNotUpdated "+result.Error.Error())
		return
	}
	var transaction models.Transaction
	transaction.Status = 1
	transaction.Settled = true
	transaction.CardId = 0
	transaction.Provider = "Grats.Cards"
	transaction.UserId = user.ID
	transaction.Amount = body.Amount
	transaction.Commission = body.Commission
	transaction.Transaction = body.Transaction
	transaction.Country = ""
	transaction.Merchant = ""
	transaction.Initial = initial
	transaction.Total = total
	transaction.Increase = true
	transaction.Balance = body.Balance
	email1, _ := c.Get("email")
	email := email1.(string)
	var user2 models.User
	h.DB.Where("email = ?", email).First(&user2)
	transaction.ManagerId = user2.ID
	h.DB.Create(&transaction)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}

func (h handler) AdminLogin(c *gin.Context) {
	body := LoginRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}

	ip := c.ClientIP()
	rToken := body.RecaptchaToken

	captcha, err := googleReCaptcha.CheckCaptcha(rToken, ip)
	if err != nil {
		printErrorMessage(c, "Captcha is invalid.")
		return
	}
	if !captcha.Success {
		printErrorMessage(c, "Captcha is invalid.")
		return
	}

	var user models.User
	if result := h.DB.Where("email = ?", body.Email).First(&user); result.Error != nil {
		printErrorMessage(c, "userNotFound")
		return
	}
	check := hashPassword.CheckPasswordHash(body.Password, user.PasswordHash)
	status := user.Status
	if !check {
		printErrorMessage(c, "incorrectPassword")
		return
	}
	if status != 2 {
		printErrorMessage(c, "registrationNotCompleted")
		return
	}
	if user.Role != 2 {
		printErrorMessage(c, "forbidden")
		return
	}
	authenticationToken, authorizationToken, err := token.AdminGetAuthenticationAuthorizationTokens(user.Email)
	if err != nil {
		printErrorMessage(c, "token not created")
		return
	}
	var session models.Session
	session.Status = 2
	session.Location = c.ClientIP() + " " + c.Request.Header["User-Agent"][0]
	session.AuthenticationToken = authenticationToken
	session.UserId = user.ID
	if result := h.DB.Save(&session); result.Error != nil {
		printErrorMessage(c, "sessionNotSet")
		return
	}
	user.LastIP = ip
	user.LastActionDate = time.Now()
	if result := h.DB.Save(&user); result.Error != nil {
		printErrorMessage(c, "userNotUpdated")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":              true,
		"user":                 &user,
		"authentication_token": authenticationToken,
		"authorization_token":  authorizationToken,
	})
}

type UserIDBody struct {
	UserID uint `json:"userID"`
}

func (h handler) DeleteUser(c *gin.Context) {
	body := UserIDBody{}
	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}
	var user models.User
	if result := h.DB.First(&user, body.UserID); result.Error != nil {
		printErrorMessage(c, "userNotFound")
		return
	}
	if result := h.DB.Delete(&user); result.Error != nil {
		printErrorMessage(c, "databaseError")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
