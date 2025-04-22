package models

import "time"

type Review struct {
	ID           uint      `json:"id"`
	SpecialistID uint      `json:"specialist_id"`
	ClientID     uint      `json:"client_id"`
	Rating       int       `json:"rating"`
	Comment      string    `json:"comment"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
