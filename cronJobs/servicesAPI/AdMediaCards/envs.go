package AdMediaCards

import (
	"github.com/spf13/viper"
)

func getDomain() string {
	return viper.Get("AD_MEDIA_CARDS_DOMAIN").(string)
}

func getAPIKey() string {
	return viper.Get("AD_MEDIA_CARDS_API_KEY").(string)
}
