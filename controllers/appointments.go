package controllers

import (
	"net/http"
	"strconv"
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
	DoctorId    uint   `json:"doctor_id" binding:"required"`
	PatientId   uint   `json:"patient_id" binding:"required"`
	Agenda      string `json:"agenda" binding:"required"`
	DateAndTime string `json:"date_time" binding:"required"`
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

func getRawAppointmentWithPatientID(id uint) models.Appointment {
	appointment := models.Appointment{}
	appointment.PatientId = id
	return appointment
}

// GET /appointments/:id
// Find a appointment
func FindAppointment(c *gin.Context) {
	// Get model if exist

	patientID, _ := strconv.ParseUint(c.Query("patient_id"), 10, 64)
	doctorID, _ := strconv.ParseUint(c.Query("doctor_id"), 10, 64)

	var appointment models.Appointment

	if c.Param("role") == "Patient" {
		if err := models.DB.Where(getRawAppointmentWithPatientID(uint(patientID))).Find(&appointment).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
	}

	if c.Param("role") == "Doctor" {
		if err := models.DB.Where(getRawAppointmentWithPatientID(uint(doctorID))).Find(&appointment).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
	}

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
	appointment := models.Appointment{DoctorId: input.DoctorId, PatientId: input.PatientId, DateAndTime: time.Now(), Agenda: input.Agenda}
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
