package db

import (
	"github.com/qrds1/Grats.Card/backend/pkg/common/models"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models/finLog"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models/taskList"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&models.Session{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&models.Card{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&models.Transaction{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&models.Supplier{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&models.Address{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&finLog.FinLog{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&finLog.CardCreation{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&finLog.CardChange{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&finLog.Replenishment{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&taskList.Task{})
	if err != nil {
		return nil
	}
	return db
}
