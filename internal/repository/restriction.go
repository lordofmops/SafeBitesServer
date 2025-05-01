package repository

import (
	"SafeBitesServer/internal/entity"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RestrictionRepository struct {
	db *gorm.DB
}

func NewRestrictionRepository(db *gorm.DB) *RestrictionRepository {
	return &RestrictionRepository{db: db}
}

func (r *RestrictionRepository) CreateRestriction(ctx context.Context, restriction *entity.Restriction) error {
	return r.db.WithContext(ctx).Create(restriction).Error
}

func (r *RestrictionRepository) GetAll(ctx context.Context) ([]*entity.Restriction, error) {
	var restrictions []*entity.Restriction
	if err := r.db.WithContext(ctx).Find(&restrictions).Error; err != nil {
		return nil, err
	}
	return restrictions, nil
}

func (r *RestrictionRepository) Add(ctx context.Context, ur *entity.UserRestriction) error {
	return r.db.WithContext(ctx).Create(ur).Error
}

func (r *RestrictionRepository) Remove(ctx context.Context, userID, restrictionID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND restriction_id = ?", userID, restrictionID).
		Delete(&entity.UserRestriction{}).Error
}

func (r *RestrictionRepository) GetUserRestrictions(ctx context.Context, userID uuid.UUID) ([]entity.Restriction, error) {
	var restrictions []entity.Restriction
	err := r.db.WithContext(ctx).
		Table("restrictions").
		Joins("JOIN user_restrictions ur ON ur.restriction_id = restrictions.restriction_id").
		Where("ur.user_id = ?", userID).
		Find(&restrictions).Error
	return restrictions, err
}
