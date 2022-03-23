package main

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/wryonik/appointment/controllers"
	"github.com/wryonik/appointment/models"

	"github.com/gin-gonic/gin"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://9b4a3071bed048bcb43bfafffbaab6e7@o1176298.ingest.sentry.io/6273811",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")

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
