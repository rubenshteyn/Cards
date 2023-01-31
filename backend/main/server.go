package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qrds1/Grats.Card/backend/cors"
	v1 "github.com/qrds1/Grats.Card/backend/pkg/api/v1"
	"github.com/qrds1/Grats.Card/backend/pkg/common/db"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("../pkg/common/envs/.env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("failed to read configs")
		fmt.Println(err)
		return
	}
	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL_PROD").(string)
	// Set the router as the default one shipped with Gin
	r := gin.Default()
	h := db.Init(dbUrl)
	r.Use(cors.Middleware())
	r.TrustedPlatform = gin.PlatformGoogleAppEngine
	// Setup route group for the API
	v1.RegisterRoutes(r, h)
	url := viper.Get("URL_DEV").(string)
	fmt.Println(url)
	err = r.RunTLS(port, "/etc/letsencrypt/live/"+url+"/fullchain.pem", "/etc/letsencrypt/live/"+url+"/privkey.pem")
	if err != nil {
		return
	}
}
