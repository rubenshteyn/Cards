package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models"
	"net/http"
)

func (h handler) AdminGetAllTransactions(c *gin.Context) {
	var transactions []models.Transaction
	if result := h.DB.Find(&transactions); result.Error != nil {
		printErrorMessage(c, "databaseError")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"transactions": &transactions,
	})
}
