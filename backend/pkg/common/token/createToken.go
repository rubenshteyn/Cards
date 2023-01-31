package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//GetAuthenticationAuthorizationTokens returns two jwt token for a user
func GetAuthenticationAuthorizationTokens(email string) (string, string, error) {
	authenticationToken, err := CreateToken(email, GetAuthenticationSecret(), time.Hour*24*30)
	if err != nil {
		return "", "", err
	}
	authorizationToken, err := CreateToken(email, GetAuthorizationSecret(), time.Minute*2)
	if err != nil {
		return "", "", err
	}
	return authenticationToken, authorizationToken, nil
}

//AdminGetAuthenticationAuthorizationTokens returns two jwt token for an admin
func AdminGetAuthenticationAuthorizationTokens(email string) (string, string, error) {
	authenticationToken, err := CreateToken(email, AdminGetAuthenticationSecret(), time.Hour*24*30)
	if err != nil {
		return "", "", err
	}
	authorizationToken, err := CreateToken(email, AdminGetAuthorizationSecret(), time.Minute*2)
	if err != nil {
		return "", "", err
	}
	return authenticationToken, authorizationToken, nil
}

func CreateToken(email string, secret string, duration time.Duration) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email"] = email
	atClaims["exp"] = time.Now().Add(duration).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
