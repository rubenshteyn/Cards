package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/qrds1/cronJobs/db"
	"github.com/qrds1/cronJobs/updateCards"
	"github.com/qrds1/cronJobs/updateTransactions"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

func runCronJobs(db *gorm.DB) {
	h := &updateCards.Handler{
		DB: db,
	}
	s := gocron.NewScheduler(time.UTC)
	//getAddresses.AddressesByJSON(h)
	_, err := s.Every(5).Minutes().Do(func() {
		fmt.Println("cron triggered 5 minutes passed")
		updateCards.StartUpdate(h)
		updateTransactions.StartUpdate(h)
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	s.StartBlocking()
}

func main() {
	viper.SetConfigFile("../backend/pkg/common/envs/.env")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	dbUrl := viper.Get("DB_URL_PROD").(string)
	h := db.Init(dbUrl)
	runCronJobs(h)
}
