package googleReCaptcha

import (
	"github.com/spf13/viper"
)

func getSiteKey() string {
	return viper.Get("GOOGLE_SITE_KEY").(string)
}

func getSecretKey() string {
	return viper.Get("GOOGLE_SECRET_KEY").(string)
}

func getDomain() string {
	return viper.Get("GOOGLE_RECAPTCHA_API").(string)
}
