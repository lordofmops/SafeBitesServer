package repository

import (
	"SafeBitesServer/internal/entity"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SearchHistoryRepository struct {
	db *gorm.DB
}

func NewSearchHistoryRepository(db *gorm.DB) *SearchHistoryRepository {
	return &SearchHistoryRepository{db: db}
}

func (r *SearchHistoryRepository) Add(ctx context.Context, h *entity.SearchHistory) error {
	var existing entity.SearchHistory
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND barcode = ?", h.UserID, h.Barcode).
		First(&existing).Error

	if err == nil {
		return r.db.WithContext(ctx).
			Model(&existing).
			Update("created_at", time.Now()).Error
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	var count int64
	if err := r.db.WithContext(ctx).
		Model(&entity.SearchHistory{}).
		Where("user_id = ?", h.UserID).
		Count(&count).Error; err != nil {
		return err
	}

	if count >= 15 {
		if err := r.db.WithContext(ctx).
			Where("user_id = ?", h.UserID).
			Order("created_at").
			Limit(1).
			Delete(&entity.SearchHistory{}).Error; err != nil {
			return err
		}
	}

	return r.db.WithContext(ctx).Create(h).Error
}

func (r *SearchHistoryRepository) GetAll(ctx context.Context, userID uuid.UUID) ([]entity.SearchHistory, error) {
	var results []entity.SearchHistory
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Find(&results).Error
	return results, err
}

func (r *SearchHistoryRepository) Delete(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&entity.SearchHistory{}).Error
}

func (r *SearchHistoryRepository) Clear(ctx context.Context, userID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&entity.SearchHistory{}).Error
}
