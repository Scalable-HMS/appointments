package models

import "time"

type Appointment struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	DoctorId    uint      `json:"doctor_id"`
	PatientId   uint      `json:"patient_id"`
	Agenda      string    `json:"agenda"`
	DateAndTime time.Time `json:"date_time"`
}
