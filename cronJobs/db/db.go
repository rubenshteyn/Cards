package db

import (
	"github.com/qrds1/cronJobs/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil
	}
	return db
}
