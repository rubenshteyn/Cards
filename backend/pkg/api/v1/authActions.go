package v1

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	email1 "github.com/qrds1/Grats.Card/backend/pkg/common/email"
	"github.com/qrds1/Grats.Card/backend/pkg/common/hashPassword"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models"
	"github.com/qrds1/Grats.Card/backend/pkg/common/servicesAPI/googleReCaptcha"
	token2 "github.com/qrds1/Grats.Card/backend/pkg/common/token"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"time"
)

type AddUserRequestBody struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Country        string `json:"country"`
	Messenger      string `json:"messenger"`
	RecaptchaToken string `json:"recaptchaToken"`
}

func (h handler) Registration(c *gin.Context) {
	body := AddUserRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}

	ip := c.ClientIP()
	token := body.RecaptchaToken

	captcha, err := googleReCaptcha.CheckCaptcha(token, ip)
	if err != nil {
		printErrorMessage(c, "Captcha is invalid.")
		return
	}
	if !captcha.Success {
		printErrorMessage(c, "Captcha is invalid.")
		return
	}

	var user models.User

	user.Name = body.Name
	user.Email = body.Email
	user.PasswordHash, _ = hashPassword.HashPassword(body.Password)
	user.Country = body.Country
	user.Messenger = body.Messenger
	user.LastIP = ip
	user.LastActionDate = time.Now()
	verificationToken, err := token2.CreateToken(body.Email, token2.GetVerificationSecret(), time.Hour)
	if err != nil {
		return
	}
	fmt.Println(verificationToken)
	emailData := email1.Data{
		URL:       "https://" + viper.Get("URL_DEV").(string) + "/verify-email/" + verificationToken,
		FirstName: user.Name,
		Subject:   "Your verification token (valid for 1 hour)",
	}
	err = email1.SendEmail("noreply@grats.cards", body.Email, &emailData, "verifyEmail.html")
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"success": false,
			"user":    "emailNotSend",
		})
		c.Abort()
		return
	}
	if result := h.DB.Create(&user); result.Error != nil {
		printErrorMessage(c, "emailAlreadyRegistered")
		return
	}
	fmt.Println("success", &user)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user":    &user,
	})
}

func (h handler) Verification(c *gin.Context) {
	verificationToken := c.Param("token")
	fmt.Println("token = ", verificationToken)
	token1, ok := checkToken(verificationToken, false, true)
	if !ok {
		printErrorMessage(c, "UnauthorizedToken")
		return
	}
	claims, ok := token1.Claims.(jwt.MapClaims)
	if !ok || !token1.Valid {
		printErrorMessage(c, "UnauthorizedToken")
		return
	}
	email, ok := claims["email"].(string)
	if !ok {
		printErrorMessage(c, "UnauthorizedEmail")
		return
	}
	var user models.User
	if result := h.DB.Where("email = ?", email).First(&user); result.Error != nil {
		printErrorMessage(c, "userNotExist")
		return
	}
	if user.Status != 0 {
		printErrorMessage(c, "emailAlreadyVerified")
		return
	}
	ip := c.ClientIP()
	user.Status = 1
	user.LastIP = ip
	user.LastActionDate = time.Now()
	if result := h.DB.Save(&user); result.Error != nil {
		printErrorMessage(c, "userNotUpdated")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
	// todo: send user to managers after email verification
}

type LoginRequestBody struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RecaptchaToken string `json:"recaptchaToken"`
}

func (h handler) Login(c *gin.Context) {
	body := LoginRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}

	ip := c.ClientIP()
	token := body.RecaptchaToken

	captcha, err := googleReCaptcha.CheckCaptcha(token, ip)
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
	authenticationToken, authorizationToken, err := token2.GetAuthenticationAuthorizationTokens(user.Email)
	if err != nil {
		printErrorMessage(c, "Tokens are not created")
		return
	}
	var session models.Session
	session.Status = 1
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

func (h handler) Logout(c *gin.Context) {
	email1, exists := c.Get("email")
	if !exists {
		printErrorMessage(c, "userNotFound")
		return
	}
	email := email1.(string)
	var user models.User
	if result := h.DB.Where("email = ?", email).First(&user); result.Error != nil {
		printErrorMessage(c, "userNotFound")
		return
	}
	header := c.Request.Header["Authorization"]
	bearerToken := strings.Split(header[0], " ")
	token := bearerToken[1]
	var session models.Session
	if result := h.DB.Delete(&session, "authentication_token = ?", token); result.Error != nil {
		printErrorMessage(c, "session not deleted")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

type GetAccessBody struct {
	Email string `json:"email"`
}

func (h handler) GetAccess(c *gin.Context) {
	body := GetAccessBody{}
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}
	var user models.User
	if result := h.DB.Where("email = ?", body.Email).First(&user); result.Error != nil {
		printErrorMessage(c, "userNotFound")
		return
	}
	authenticationToken, _, _ := token2.GetAuthenticationAuthorizationTokens(user.Email)
	var session models.Session
	session.Status = 4
	session.Location = c.ClientIP() + " " + c.Request.Header["User-Agent"][0]
	session.AuthenticationToken = authenticationToken
	session.UserId = user.ID
	if result := h.DB.Save(&session); result.Error != nil {
		printErrorMessage(c, "sessionNotSet")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   authenticationToken,
	})
}

func (h handler) Access(c *gin.Context) {
	email1, exists := c.Get("email")
	if !exists {
		printErrorMessage(c, "Cannot check email")
		return
	}
	email := email1.(string)
	header := c.Request.Header["Authorization"]
	bearerToken := strings.Split(header[0], " ")
	token := bearerToken[1]
	var user models.User
	if result := h.DB.Where("email = ?", email).First(&user); result.Error != nil {
		printErrorMessage(c, "userNotFound")
		return
	}
	authorizationToken, err := token2.CreateToken(email, token2.GetAuthorizationSecret(), time.Minute*2)
	if err != nil {
		printErrorMessage(c, "Cannot create a token")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":              true,
		"user":                 &user,
		"authorization_token":  authorizationToken,
		"authentication_token": token,
	})
}

type ForgotPassword struct {
	Email string `json:"email"`
}

func (h handler) ForgotPassword(c *gin.Context) {
	body := ForgotPassword{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}

	var user models.User
	if result := h.DB.Where("email = ?", body.Email).First(&user); result.Error != nil {
		printErrorMessage(c, "userNotExist")
		return
	}
	resetToken, err := token2.CreateToken(body.Email, token2.GetResetPasswordSecret(), time.Hour)
	if err != nil {
		return
	}
	fmt.Println(resetToken)
	// todo: mail token with get param
	emailData := email1.Data{
		URL:       "https://test.grats.cards/reset-password/" + resetToken,
		FirstName: user.Name,
		Subject:   "Your reset password token (valid for 1 hour)",
	}
	err = email1.SendEmail("noreply@grats.cards", body.Email, &emailData, "resetPassword.html")
	if err != nil {
		printErrorMessage(c, "emailNotSend")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}

type ResetPassword struct {
	NewPassword string `json:"newPassword"`
}

func (h handler) ResetPassword(c *gin.Context) {
	body := ResetPassword{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}
	resetToken := c.Param("token")
	fmt.Println("token = ", resetToken)
	token1, ok := checkToken(resetToken, false, false)
	if !ok {
		printErrorMessage(c, "UnauthorizedToken")
		return
	}
	claims, ok := token1.Claims.(jwt.MapClaims)
	if !ok || !token1.Valid {
		printErrorMessage(c, "UnauthorizedToken")
		return
	}
	email, ok := claims["email"].(string)
	if !ok {
		printErrorMessage(c, "UnauthorizedEmail")
		return
	}
	var user models.User
	if result := h.DB.Where("email = ?", email).First(&user); result.Error != nil {
		printErrorMessage(c, "UserNotExist")
		return
	}
	user.PasswordHash, _ = hashPassword.HashPassword(body.NewPassword)
	ip := c.ClientIP()
	user.LastIP = ip
	user.LastActionDate = time.Now()
	if result := h.DB.Save(&user); result.Error != nil {
		printErrorMessage(c, "UserNotUpdated")
		return
	}
	// deactivate all sessions
	var session models.Session
	if result := h.DB.Delete(&session, "user_id = ?", user.ID); result.Error != nil {
		printErrorMessage(c, "databaseError")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
