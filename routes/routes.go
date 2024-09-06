package routes

import (
	"github.com/cauelz/golang-event-booking-rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// The POST, PUT, and DELETE routes are protected by the Authenticate middleware.
	// Created a Group called / that uses the Authenticate middleware.
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerEvent)
	authenticated.DELETE("/events/:id/register", unregisterEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}