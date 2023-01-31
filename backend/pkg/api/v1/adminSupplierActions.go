package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models"
	"net/http"
)

func (h handler) GetSuppliers(c *gin.Context) {
	var suppliers []models.Supplier
	if result := h.DB.Find(&suppliers); result.Error != nil {
		printErrorMessage(c, "databaseError")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"suppliers": &suppliers,
	})
}

func (h handler) GetBins(c *gin.Context) {
	var supplier models.Supplier
	if result := h.DB.Find(&supplier, 1); result.Error != nil {
		printErrorMessage(c, "databaseError")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"bins":    supplier.Bins,
	})
}

type UpdateSupplierBody struct {
	ID     uint   `json:"id"`
	Status int    `json:"status"`
	Bins   string `json:"bins"`
}

func (h handler) UpdateSupplier(c *gin.Context) {
	body := UpdateSupplierBody{}
	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		printErrorMessage(c, "badRequest")
		return
	}

	var supplier models.Supplier
	h.DB.First(&supplier, body.ID)
	if body.Status != supplier.Status {
		supplier.Status = body.Status
	}
	if body.Bins != supplier.Bins {
		supplier.Bins = body.Bins
	}
	if result := h.DB.Save(&supplier); result.Error != nil {
		printErrorMessage(c, "supplierNotUpdated "+result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
