package middlewares

import (
	"net/http"

	"github.com/cauelz/golang-event-booking-rest-api/utils"
	"github.com/gin-gonic/gin"
)


func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userId, error := utils.VerifyToken(token)

	if error != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.Set("userId", userId)
	c.Next()
}