package updateTransactions

import (
	"fmt"
	"github.com/qrds1/cronJobs/models"
	"github.com/qrds1/cronJobs/servicesAPI/AdMediaCards"
	"github.com/qrds1/cronJobs/updateCards"
	"time"
)

func StartUpdate(h *updateCards.Handler) {
	r, err := AdMediaCards.GetTransactions(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	addNewTransactions(r, h)
}

func addNewTransactions(response *AdMediaCards.GetTransactionsResponse, h *updateCards.Handler) {
	transactions := response.Result
	for _, transaction := range transactions {
		var ourTransaction models.Transaction
		var create bool = true
		if result := h.DB.Where("provider_id = ? AND provider = ?", transaction.TransactionID, "AdMediaCards").First(&ourTransaction); result.Error == nil {
			create = false
			continue
		}
		ourTransaction.Merchant = transaction.Merchant
		ourTransaction.Amount = transaction.Amount
		ourTransaction.TransactionDate = time.Unix(transaction.TransactionDate, 0)
		ourTransaction.Provider = "AdMediaCards"
		ourTransaction.ProviderId = transaction.TransactionID
		ourTransaction.CardId = transaction.CardID
		ourTransaction.Settled = !transaction.IsTemp
		var card models.Card
		if result := h.DB.Where("provider_id = ? AND card_provider = ?", transaction.CardID, "AdMediaCards").First(&card); result.Error == nil {
			ourTransaction.UserId = card.UserId
			ourTransaction.CardId = card.ID
		}
		ourTransaction.Amount = transaction.Amount
		ourTransaction.Country = transaction.Country
		ourTransaction.Merchant = transaction.Merchant
		ourTransaction.Increase = false
		if create {
			if result := h.DB.Create(&ourTransaction); result.Error != nil {
				fmt.Println("card not updated " + result.Error.Error())
			}
			continue
		}
		h.DB.Save(&ourTransaction)
	}
}
