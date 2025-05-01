package entity

import "github.com/google/uuid"

type Restriction struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name string    `gorm:"not null" json:"name"`
	Type string    `gorm:"not null" json:"type"`
	Tag  string    `gorm:"not null" json:"tag"`
}

type UserRestriction struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UserID        uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	RestrictionID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
}
