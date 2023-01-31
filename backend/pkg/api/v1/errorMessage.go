package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func printErrorMessage(c *gin.Context, errorMessage string) {
	fmt.Println(errorMessage)
	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"error":   errorMessage,
	})
	c.Abort()
}
