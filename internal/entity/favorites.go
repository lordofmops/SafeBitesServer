package entity

import "github.com/google/uuid"

type Favorites struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UserID  uuid.UUID `gorm:"not null" json:"-"`
	Barcode string    `gorm:"not null" json:"barcode"`
}
