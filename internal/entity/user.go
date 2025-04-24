package entity

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Login        string    `gorm:"unique" json:"login"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Name         string    `json:"name"`
}
