package usecase

import (
	"SafeBitesServer/internal/entity"
	"context"

	"github.com/google/uuid"
)

type SearchHistoryRepository interface {
	Add(context.Context, *entity.SearchHistory) error
	GetAll(context.Context, uuid.UUID) ([]entity.SearchHistory, error)
	Delete(context.Context, uuid.UUID, uuid.UUID) error
	Clear(context.Context, uuid.UUID) error
}

type SearchHistoryUsecase struct {
	repo SearchHistoryRepository
}

func NewSearchHistoryUsecase(r SearchHistoryRepository) *SearchHistoryUsecase {
	return &SearchHistoryUsecase{repo: r}
}

func (uc *SearchHistoryUsecase) Add(ctx context.Context, userID uuid.UUID, barcode string) error {
	return uc.repo.Add(ctx, &entity.SearchHistory{
		UserID:  userID,
		Barcode: barcode,
	})
}

func (uc *SearchHistoryUsecase) GetAll(ctx context.Context, userID uuid.UUID) ([]entity.SearchHistory, error) {
	return uc.repo.GetAll(ctx, userID)
}

func (uc *SearchHistoryUsecase) Delete(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	return uc.repo.Delete(ctx, id, userID)
}

func (uc *SearchHistoryUsecase) Clear(ctx context.Context, userID uuid.UUID) error {
	return uc.repo.Clear(ctx, userID)
}
