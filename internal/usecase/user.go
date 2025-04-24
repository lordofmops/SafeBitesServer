package usecase

import (
	"SafeBitesServer/internal/entity"
	"context"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetAll(context.Context) ([]*entity.User, error)
	GetByID(context.Context, uuid.UUID) (*entity.User, error)
	GetByLogin(context.Context, string) (*entity.User, error)
	UpdateName(context.Context, uuid.UUID, string) error
	Delete(context.Context, uuid.UUID) error
}

type UserUsecase struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) GetAll(ctx context.Context) ([]*entity.User, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *UserUsecase) GetProfile(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *UserUsecase) UpdateName(ctx context.Context, id uuid.UUID, name string) error {
	return uc.repo.UpdateName(ctx, id, name)
}

func (uc *UserUsecase) DeleteProfile(ctx context.Context, id uuid.UUID) error {
	return uc.repo.Delete(ctx, id)
}
