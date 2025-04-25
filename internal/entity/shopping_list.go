package entity

import "github.com/google/uuid"

type ShoppingList struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Name   string    `json:"name"`
}

type ShoppingListProduct struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	ShoppingListID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Barcode        string    `gorm:"not null" json:"barcode"`
}
