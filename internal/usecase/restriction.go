package usecase

import (
	"SafeBitesServer/internal/entity"
	"context"
	"github.com/google/uuid"
)

type RestrictionRepo interface {
	CreateRestriction(context.Context, *entity.Restriction) error
	Add(context.Context, *entity.UserRestriction) ([]entity.Restriction, error)
	Remove(context.Context, uuid.UUID, uuid.UUID) ([]entity.Restriction, error)
	GetUserRestrictions(context.Context, uuid.UUID) ([]entity.Restriction, error)
	GetAll(context.Context) ([]*entity.Restriction, error)
}

type RestrictionUsecase struct {
	repo RestrictionRepo
}

func NewRestrictionUsecase(r RestrictionRepo) *RestrictionUsecase {
	return &RestrictionUsecase{repo: r}
}

func (uc *RestrictionUsecase) CreateRestriction(ctx context.Context, name string, _type string, tag string) error {
	restriction := &entity.Restriction{
		Name: name,
		Type: _type,
		Tag:  tag,
	}
	return uc.repo.CreateRestriction(ctx, restriction)
}

func (uc *RestrictionUsecase) GetAll(ctx context.Context) ([]*entity.Restriction, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *RestrictionUsecase) Add(ctx context.Context, userID, restrictionID uuid.UUID) ([]entity.Restriction, error) {
	return uc.repo.Add(ctx, &entity.UserRestriction{
		UserID:        userID,
		RestrictionID: restrictionID,
	})
}

func (uc *RestrictionUsecase) Remove(ctx context.Context, userID, restrictionID uuid.UUID) ([]entity.Restriction, error) {
	return uc.repo.Remove(ctx, userID, restrictionID)
}

func (uc *RestrictionUsecase) GetUserRestrictions(ctx context.Context, userID uuid.UUID) ([]entity.Restriction, error) {
	return uc.repo.GetUserRestrictions(ctx, userID)
}
