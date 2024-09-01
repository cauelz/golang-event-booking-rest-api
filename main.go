package main

import (
	"github.com/cauelz/golang-event-booking-rest-api/db"
	"github.com/cauelz/golang-event-booking-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
