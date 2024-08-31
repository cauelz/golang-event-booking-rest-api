package main

import (
	"net/http"

	"github.com/cauelz/golang-event-booking-rest-api/db"
	"github.com/cauelz/golang-event-booking-rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(c *gin.Context) {

	events, error := models.GetAllEvents()

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get events"})
		return
	}

	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {

	var event models.Event

	err := c.ShouldBindBodyWithJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind JSON"})
		return
	}

	event.ID = 1
	event.UserId = 1

	event.Save()

	c.JSON(http.StatusCreated, gin.H{"status": "Event created successfully", "event": event})
}
