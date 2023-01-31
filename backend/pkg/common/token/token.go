package token

import (
	"github.com/spf13/viper"
)

func GetAuthorizationSecret() string {
	secret := viper.Get("ACCESS_SECRET_AUTHORIZATION").(string)
	if secret == "" {
		secret = "dasjdksahdjsahdkash"
	}
	return secret
}

func GetAuthenticationSecret() string {
	secret := viper.Get("ACCESS_SECRET_AUTHENTICATION").(string)
	if secret == "" {
		secret = "fdsjhfdsjkhflakjekwanejah"
	}
	return secret
}

func GetVerificationSecret() string {
	secret := viper.Get("ACCESS_SECRET_VERIFICATION").(string)
	if secret == "" {
		secret = "fhsgfjhasabdnmsavdmanb"
	}
	return secret
}

func GetResetPasswordSecret() string {
	secret := viper.Get("ACCESS_SECRET_RESET_PASSWORD").(string)
	if secret == "" {
		secret = "fhsgfjhasabdnmsavdmanb"
	}
	return secret
}

func AdminGetAuthenticationSecret() string {
	secret := viper.Get("ACCESS_SECRET_AUTHENTICATION_ADMIN").(string)
	if secret == "" {
		secret = "fhsgfjhasabdnmsavdmanb"
	}
	return secret
}

func AdminGetAuthorizationSecret() string {
	secret := viper.Get("ACCESS_SECRET_AUTHORIZATION_ADMIN").(string)
	if secret == "" {
		secret = "fhsgfjhasabdnmsavdmanb"
	}
	return secret
}
