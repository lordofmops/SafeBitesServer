package repository

import (
	"SafeBitesServer/internal/entity"
	"context"

	"gorm.io/gorm"
)

type StoreRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *StoreRepository {
	return &StoreRepository{db: db}
}

func (r *StoreRepository) GetAll(ctx context.Context) ([]*entity.Store, error) {
	var stores []*entity.Store
	if err := r.db.WithContext(ctx).Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}

func (r *StoreRepository) CreateStore(ctx context.Context, store *entity.Store) error {
	return r.db.WithContext(ctx).Create(store).Error
}
