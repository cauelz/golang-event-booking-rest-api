package routes

import (
	"net/http"
	"strconv"

	"github.com/cauelz/golang-event-booking-rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {

	events, error := models.GetAllEvents()

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get events"})
		return
	}

	c.JSON(http.StatusOK, events)
}

func getEvent(c *gin.Context) {

	id, error := strconv.ParseInt(c.Param("id"), 10, 64)

	if(error != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	event, error := models.GetEventById(id)

	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {

	var event models.Event

	err := c.ShouldBindBodyWithJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind JSON"})
		return
	}

	userId := c.GetInt64("userId")

	event.UserId = userId

	event.Save()

	c.JSON(http.StatusCreated, gin.H{"status": "Event created successfully", "event": event})
}

func updateEvent(c *gin.Context) {
	id, error := strconv.ParseInt(c.Param("id"), 10, 64)

	if(error != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	userId := c.GetInt64("userId")
	event, error := models.GetEventById(id)

	
	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	
	if userId != event.UserId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this event"})
		return
	}
	var updatedEvent models.Event

	error = c.ShouldBindBodyWithJSON(&updatedEvent)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind JSON"})
		return
	}

	updatedEvent.ID = id

	error = updatedEvent.Update()

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Event updated successfully", "event": updatedEvent})

}

func deleteEvent(c *gin.Context) {
	id, error := strconv.ParseInt(c.Param("id"), 10, 64)

	if(error != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	userId := c.GetInt64("userId")
	event, error := models.GetEventById(id)

	
	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if userId != event.UserId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this event"})
		return
	}

	error = models.DeleteEvent(id)

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Event deleted successfully"})

}