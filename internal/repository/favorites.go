package repository

import (
	"SafeBitesServer/internal/entity"
	"context"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type FavoritesRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) *FavoritesRepository {
	return &FavoritesRepository{db: db}
}

func (r *FavoritesRepository) Add(ctx context.Context, userID uuid.UUID, barcode string) error {
	var exists bool
	err := r.db.WithContext(ctx).
		Model(&entity.Favorites{}).
		Select("count(*) > 0").
		Where("user_id = ? AND barcode = ?", userID, barcode).
		Find(&exists).Error
	if err != nil {
		return err
	}
	if exists {
		return nil // уже в избранном — не добавляем
	}
	return r.db.WithContext(ctx).Create(&entity.Favorites{
		UserID:  userID,
		Barcode: barcode,
	}).Error
}

func (r *FavoritesRepository) Delete(ctx context.Context, userID uuid.UUID, barcode string) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND barcode = ?", userID, barcode).
		Delete(&entity.Favorites{}).Error
}

func (r *FavoritesRepository) List(ctx context.Context, userID uuid.UUID) ([]string, error) {
	var barcodes []string
	err := r.db.WithContext(ctx).
		Model(&entity.Favorites{}).
		Where("user_id = ?", userID).
		Pluck("barcode", &barcodes).Error
	return barcodes, err
}
