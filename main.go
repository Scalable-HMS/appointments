package main

import (
	"log"
	"time"
		"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"


	"github.com/getsentry/sentry-go"
	"github.com/wryonik/appointment/controllers"
	"github.com/wryonik/appointment/models"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Role  string `json:"given_name"`
	Email string `json:"email"`
	Id    string `json:"nickname"`
}

func authMid(c *gin.Context) {

	url := "https://dev-rgmfg73e.us.auth0.com/userinfo"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", c.Request.Header["Authorization"][0])

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := Response{}
	json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Email)
	fmt.Println(response.Role)
	fmt.Println(response.Id)
	c.Params = []gin.Param{
		{
			Key:   "email",
			Value: response.Email,
		},
		{
			Key:   "role",
			Value: response.Role,
		},
		{
			Key:   "id",
			Value: response.Id,
		},
		
	}
}

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

	secureGroup := r.Group("/secure/", authMid)

	// Routes
	secureGroup.GET("/appointment", controllers.FindAppointments)
	secureGroup.POST("/appointment", controllers.CreateAppointment)
	secureGroup.PATCH("/appointment", controllers.UpdateAppointment)
	secureGroup.DELETE("/appointment", controllers.DeleteAppointment)

	// Run the server
	r.Run()

}
