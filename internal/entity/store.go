package entity

import "github.com/google/uuid"

type Store struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name string    `gorm:"not null" json:"name"`
	Link string    `gorm:"not null" json:"link"`
}
