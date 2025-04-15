package models

type Appointment struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	ClientID     uint   `json:"client_id"`
	SpecialistID uint   `json:"specialist_id"`
	Datetime     string `json:"datetime"`
	Notes        string `json:"notes"`
}
