package entity

import "github.com/google/uuid"

type History struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID uuid.UUID `gorm:"not null"`
}

type HistoryProduct struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	HistoryID uuid.UUID `gorm:"not null"`
	Barcode   string    `gorm:"not null"`
}
