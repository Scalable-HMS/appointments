package main

import (
	"github.com/wryonik/appointment/controllers"
	"github.com/wryonik/appointment/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/appointment", controllers.FindAppointments)
	r.POST("/appointment", controllers.CreateAppointment)
	r.PATCH("/appointment", controllers.UpdateAppointment)
	r.DELETE("/appointment", controllers.DeleteAppointment)

	// Run the server
	r.Run()
}
