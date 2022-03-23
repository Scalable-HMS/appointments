package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wryonik/appointment/models"
)

type Appointment struct {
	DoctorId    uint      `json:"doctor_id"`
	PatientId   uint      `json:"patient_id"`
	Agenda      string    `json:"agenda"`
	DateAndTime time.Time `json:"date_time"`
}

type CreateAppointmentInput struct {
	DoctorId    uint      `json:"doctor_id" binding:"required"`
	PatientId   uint      `json:"patient_id" binding:"required"`
	Agenda      string    `json:"agenda" binding:"required"`
	DateAndTime time.Time `json:"date_time" binding:"required"`
}

type UpdateAppointmentInput struct {
	DoctorId    uint      `json:"doctor_id""`
	PatientId   uint      `json:"patient_id""`
	Agenda      string    `json:"agenda""`
	DateAndTime time.Time `json:"date_time""`
}

// GET /appointments
// Find all appointments
func FindAppointments(c *gin.Context) {
	var appointments []models.Appointment
	models.DB.Find(&appointments)

	c.JSON(http.StatusOK, gin.H{"data": appointments})
}

// GET /appointments/:id
// Find a appointment
func FindAppointment(c *gin.Context) {
	// Get model if exist
	var appointment models.Appointment
	if err := models.DB.Where("id = ?", c.Param("id")).First(&appointment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": appointment})
}

// POST /appointments
// Create new appointment
func CreateAppointment(c *gin.Context) {
	// Validate input
	var input CreateAppointmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create appointment
	appointment := models.Appointment{DoctorId: input.DoctorId, PatientId: input.PatientId, DateAndTime: input.DateAndTime, Agenda: input.Agenda}
	models.DB.Create(&appointment)

	c.JSON(http.StatusOK, gin.H{"data": appointment})
}

// PATCH /appointments/:id
// Update a appointment
func UpdateAppointment(c *gin.Context) {
	// Get model if exist
	var appointment models.Appointment
	if err := models.DB.Where("id = ?", c.Param("id")).First(&appointment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateAppointmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&appointment).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": appointment})
}

// DELETE /appointments/:id
// Delete a appointment
func DeleteAppointment(c *gin.Context) {
	// Get model if exist
	var appointment models.Appointment
	if err := models.DB.Where("id = ?", c.Param("id")).First(&appointment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&appointment)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
