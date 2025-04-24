package entity

import "github.com/google/uuid"

type Store struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name string    `gorm:"not null"`
	Tag  string    `gorm:"not null"`
	Link string    `gorm:"not null"`
}
