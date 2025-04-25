package usecase

import (
	"SafeBitesServer/internal/entity"
	"context"
)

type StoreRepository interface {
	GetAll(context.Context) ([]*entity.Store, error)
	CreateStore(context.Context, *entity.Store) error
}
type StoreUsecase struct {
	repo StoreRepository
}

func NewStoreUsecase(repo StoreRepository) *StoreUsecase {
	return &StoreUsecase{repo: repo}
}

func (uc *StoreUsecase) CreateStore(ctx context.Context, name string, link string) error {
	list := &entity.Store{
		Name: name,
		Link: link,
	}
	return uc.repo.CreateStore(ctx, list)
}

func (uc *StoreUsecase) GetAll(ctx context.Context) ([]*entity.Store, error) {
	return uc.repo.GetAll(ctx)
}
