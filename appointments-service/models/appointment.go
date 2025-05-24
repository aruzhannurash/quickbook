package models

import "gorm.io/gorm"

type Appointment struct {
	gorm.Model
	ClientID     uint   `json:"client_id"`
	SpecialistID uint   `json:"specialist_id"`
	Datetime     string `json:"datetime"`
}
