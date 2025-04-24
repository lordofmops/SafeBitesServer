package entity

import "github.com/google/uuid"

type Restrictions struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	RestrictionID uuid.UUID `gorm:"not null"`
	Name          string    `gorm:"not null"`
	Type          string    `gorm:"not null"`
	Tag           string    `gorm:"not null"`
}

type UserRestriction struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID uuid.UUID `gorm:"not null"`
}
