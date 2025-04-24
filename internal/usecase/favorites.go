package usecase

import (
	"SafeBitesServer/internal/entity"
	"context"

	"github.com/google/uuid"
)

type FavoritesUsecase struct {
	repo FavoriteRepo
}

type FavoriteRepo interface {
	Add(context.Context, *entity.Favorites) error
	Delete(context.Context, uuid.UUID, string) error
	List(context.Context, uuid.UUID) ([]string, error)
}

func NewFavoritesUsecase(r FavoriteRepo) *FavoritesUsecase {
	return &FavoritesUsecase{repo: r}
}

func (uc *FavoritesUsecase) Add(ctx context.Context, userID uuid.UUID, barcode string) error {
	return uc.repo.Add(ctx, &entity.Favorites{
		UserID:  userID,
		Barcode: barcode,
	})
}

func (uc *FavoritesUsecase) Delete(ctx context.Context, userID uuid.UUID, barcode string) error {
	return uc.repo.Delete(ctx, userID, barcode)
}

func (uc *FavoritesUsecase) List(ctx context.Context, userID uuid.UUID) ([]string, error) {
	return uc.repo.List(ctx, userID)
}
