package routes

import (
	"net/http"
	"strconv"

	"github.com/cauelz/golang-event-booking-rest-api/models"
	"github.com/gin-gonic/gin"
)


func registerEvent(c *gin.Context) {

	userId := c.GetInt64("userId")
	eventId, error := strconv.ParseInt(c.Param("id"), 10, 64)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse event ID"})
		return
	}

	event, error :=models.GetEventById(eventId)

	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	error = event.Register(userId)

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not register user to event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered to event"})

}

func unregisterEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, error := strconv.ParseInt(c.Param("id"), 10, 64)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse event ID"})
		return
	}

	error = models.Unregister(userId, eventId)

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not unregister user to event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user unregistered to event"})

}