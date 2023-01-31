package tgBot

import (
	"fmt"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models"
	"github.com/spf13/viper"
)

func SendRequest(user models.User) {
	bot, err := tgBotApi.NewBotAPI(viper.Get("TELEGRAM_APITOKEN").(string))
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	text := fmt.Sprintf("Запрос от %s(ID: %d)", user.Name, user.ID)
	//var chatID int64 = 5645928896
	var chatID int64 = 339820425
	msg := tgBotApi.NewMessage(chatID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
