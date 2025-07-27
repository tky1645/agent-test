package entities

import (
	"time"
)

type WateringRecord struct {
	ID        string     `json:"id"`
	PlantID   string     `json:"plant_id"`
	WateredAt time.Time  `json:"watered_at"`
	Notes     *string    `json:"notes"`
	CreatedAt time.Time  `json:"created_at"`
}
