package entity

import (
	"github.com/google/uuid"
	"time"
)

type SearchHistory struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Barcode   string    `gorm:"not null" json:"barcode"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
}
