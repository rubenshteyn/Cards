package v1

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models"
	token2 "github.com/qrds1/Grats.Card/backend/pkg/common/token"
	"net/http"
	"strings"
	"time"
)

func (h handler) AdminCheckTokenAuthHandler(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header["Authorization"]
		bearerToken := strings.Split(header[0], " ")
		if len(bearerToken) != 2 {
			printErrorMessage(c, "No token found")
			return
		}
		if bearerToken[0] != "Bearer" {
			printErrorMessage(c, "Error in authorization token. it needs to be in form of 'Bearer <token>'")
			return
		}
		token, ok := AdminCheckToken(bearerToken[1], false)
		if !ok {
			printErrorMessage(c, "UnauthorizedToken")
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			email, ok := claims["email"].(string)
			if !ok {
				printErrorMessage(c, "UnauthorizedEmail")
				return
			}
			var user models.User
			if result := h.DB.Where("email = ?", email).First(&user); result.Error != nil {
				printErrorMessage(c, "Unauthorized, user not exists")
				return
			}
			if user.Role != 2 {
				printErrorMessage(c, "accessForbidden")
				return
			}
			var session models.Session
			if result := h.DB.Where("user_id = ? AND authentication_token = ?", user.ID, bearerToken[1]).First(&session); result.Error != nil {
				printErrorMessage(c, "unauthorizedToken")
				return
			}
			if session.Status == 0 || session.Status == 3 {
				printErrorMessage(c, "tokenDisabled")
				return
			}
			c.Set("email", email)
		}
		next(c)
	}
}

func (h handler) AdminCheckTokenHandler(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header["Authorization"]
		bearerToken := strings.Split(header[0], " ")
		if len(bearerToken) != 2 {
			printErrorMessage(c, "No token found")
			return
		}
		if bearerToken[0] != "Bearer" {
			printErrorMessage(c, "Error in authorization token. it needs to be in form of 'Bearer <token>'")
			return
		}
		token, ok := AdminCheckToken(bearerToken[1], true)
		if !ok {
			printErrorMessage(c, "UnauthorizedToken")
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			email, ok := claims["email"].(string)
			if !ok {
				printErrorMessage(c, "UnauthorizedEmail")
				return
			}
			var user models.User
			if result := h.DB.Where("email = ?", email).First(&user); result.Error != nil {
				printErrorMessage(c, "Unauthorized, user not exists")
				return
			}
			if user.Role != 2 {
				printErrorMessage(c, "accessForbidden")
				return
			}
			c.Set("email", email)
		}
		next(c)
	}
}

func AdminCheckToken(tokenString string, authorizationToken bool) (*jwt.Token, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		if authorizationToken {
			return []byte(token2.AdminGetAuthorizationSecret()), nil
		} else {
			return []byte(token2.AdminGetAuthenticationSecret()), nil
		}
	})
	if err != nil {
		return nil, false
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, false
	}
	return token, true
}

func AdminGetTokenByToken(c *gin.Context) {
	email1, exists := c.Get("email")
	if !exists {
		printErrorMessage(c, "Cannot check email")
		return
	}
	email := email1.(string)
	authorizationToken, err := token2.CreateToken(email, token2.AdminGetAuthorizationSecret(), time.Minute*2)
	if err != nil {
		printErrorMessage(c, "Cannot create a token")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":             true,
		"authorization_token": authorizationToken,
	})
}

func (h handler) CheckTokenAuthHandler(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header["Authorization"]
		bearerToken := strings.Split(header[0], " ")
		if len(bearerToken) != 2 {
			printErrorMessage(c, "No token found")
			return
		}
		if bearerToken[0] != "Bearer" {
			printErrorMessage(c, "Error in authorization token. it needs to be in form of 'Bearer <token>'")
			return
		}
		token, ok := checkToken(bearerToken[1], false)
		if !ok {
			printErrorMessage(c, "UnauthorizedToken")
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			email, ok := claims["email"].(string)
			if !ok {
				printErrorMessage(c, "UnauthorizedEmail")
				return
			}
			var user models.User
			if result := h.DB.Where("email = ?", email).First(&user); result.Error != nil {
				printErrorMessage(c, "Unauthorized, user not exists")
				return
			}
			var session models.Session
			if result := h.DB.Where("user_id = ? AND authentication_token = ?", user.ID, bearerToken[1]).First(&session); result.Error != nil {
				printErrorMessage(c, "UnauthorizedToken")
				return
			}
			if session.Status == 0 {
				printErrorMessage(c, "tokenDisabled")
				return
			}
			c.Set("email", email)
		}
		next(c)
	}
}

func checkToken(tokenString string, authorizationToken bool, verificationToken ...bool) (*jwt.Token, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		if len(verificationToken) > 0 {
			if !verificationToken[0] {
				return []byte(token2.GetResetPasswordSecret()), nil
			} else {
				return []byte(token2.GetVerificationSecret()), nil
			}
		}
		if authorizationToken {
			return []byte(token2.GetAuthorizationSecret()), nil
		} else {
			return []byte(token2.GetAuthenticationSecret()), nil
		}
	})
	if err != nil {
		return nil, false
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, false
	}
	return token, true
}

func GetTokenByToken(c *gin.Context) {
	email1, exists := c.Get("email")
	if !exists {
		printErrorMessage(c, "Cannot check Email")
		return
	}
	email := email1.(string)
	authorizationToken, err := token2.CreateToken(email, token2.GetAuthorizationSecret(), time.Minute*2)
	if err != nil {
		printErrorMessage(c, "Cannot create a token")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":             true,
		"authorization_token": authorizationToken,
	})
}

func (h handler) CheckTokenHandler(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header["Authorization"]
		bearerToken := strings.Split(header[0], " ")
		if len(bearerToken) != 2 {
			printErrorMessage(c, "No token found")
			return
		}
		if bearerToken[0] != "Bearer" {
			printErrorMessage(c, "Error in authorization token. it needs to be in form of 'Bearer <token>'")
			return
		}
		token, ok := checkToken(bearerToken[1], true)
		if !ok {
			printErrorMessage(c, "UnauthorizedToken")
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			email, ok := claims["email"].(string)
			if !ok {
				printErrorMessage(c, "UnauthorizedEmail")
				return
			}
			var user models.User
			if result := h.DB.Where("email = ?", email).First(&user); result.Error != nil {
				printErrorMessage(c, "Unauthorized, user not exists")
				return
			}
			c.Set("email", email)
		}
		next(c)
	}
}
