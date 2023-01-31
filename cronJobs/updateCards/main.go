package updateCards

import (
	"encoding/json"
	"fmt"
	"github.com/qrds1/cronJobs/models"
	"github.com/qrds1/cronJobs/servicesAPI/AdMediaCards"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func StartUpdate(h *Handler) {
	r, err := AdMediaCards.GetCards(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	cards := r.Result
	for _, card := range cards {
		var ourCard models.Card
		if result := h.DB.Where("provider_id = ? AND card_provider = ?", card.CardID, "AdMediaCards").First(&ourCard); result.Error != nil {
			a, _ := json.Marshal(card)
			fmt.Println("card not found in our system")
			fmt.Println(string(a))
			continue
		}
		resp, err := AdMediaCards.GetCardCVV(card.CardID)
		if err != nil {
			fmt.Println(err)
			return
		}
		cvv := resp.Result
		ourCard.Limit = card.Limit
		ourCard.Balance = card.Balance
		ourCard.Spend = card.Spend
		ourCard.Name = card.Name
		ourCard.Address = card.Address
		ourCard.CardNumber = cvv.CcNum
		ourCard.CVV = cvv.Cvx2
		ourCard.ExpirationDate = card.ExpMonth + "/" + card.ExpYear
		ourCard.Status = 1
		if !card.IsActive {
			ourCard.Status = 0
		}
		if result := h.DB.Save(&ourCard); result.Error != nil {
			fmt.Println("card not updated " + result.Error.Error())
			return
		}
	}
}
