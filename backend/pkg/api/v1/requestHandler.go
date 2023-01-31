package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qrds1/Grats.Card/backend/pkg/common/models"
	"strings"
	"time"
)

func (h handler) RequestHandler(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := strings.Split(c.Request.Header["Authorization"][0], " ")
		ip := c.ClientIP()
		if len(tokenHeader) != 2 {
			next(c)
			return
		}
		token := tokenHeader[1]
		var session models.Session
		if result := h.DB.Where("authentication_token = ?", token).First(&session); result.Error != nil {
			return
		}
		if session.Status == 4 {
			next(c)
		}
		var user models.User
		if result := h.DB.First(&user, session.UserId); result.Error != nil {
			next(c)
		}
		user.LastIP = ip
		user.LastActionDate = time.Now()
		h.DB.Save(&user)
		next(c)
	}
}
