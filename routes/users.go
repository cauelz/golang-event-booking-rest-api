package routes

import (
	"net/http"

	"github.com/cauelz/golang-event-booking-rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup (c *gin.Context) {

	var user models.User

	error := c.ShouldBindJSON(&user)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Could not bind JSON"})
		return
	}

	error = user.Save()

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "User created", "user": user})
}