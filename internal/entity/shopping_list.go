package entity

import "github.com/google/uuid"

type ShoppingList struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID uuid.UUID `gorm:"not null"`
}

type ShoppingListProduct struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ShoppingListID uuid.UUID `gorm:"not null"`
	Barcode        string    `gorm:"not null"`
}
