package usecase

import (
	"SafeBitesServer/internal/entity"
	"context"

	"github.com/google/uuid"
)

type ShoppingListRepository interface {
	CreateList(context.Context, *entity.ShoppingList) error
	GetLists(context.Context, uuid.UUID) ([]entity.ShoppingList, error)
	DeleteList(context.Context, uuid.UUID, uuid.UUID) error
	UpdateListName(context.Context, uuid.UUID, string) error
	AddProduct(context.Context, *entity.ShoppingListProduct) error
	DeleteProduct(context.Context, uuid.UUID, uuid.UUID) error
	GetProducts(context.Context, uuid.UUID, uuid.UUID) ([]entity.ShoppingListProduct, error)
}

type ShoppingListUsecase struct {
	repo ShoppingListRepository
}

func NewShoppingListUsecase(r ShoppingListRepository) *ShoppingListUsecase {
	return &ShoppingListUsecase{repo: r}
}

func (uc *ShoppingListUsecase) CreateList(ctx context.Context, userID uuid.UUID, name string) error {
	list := &entity.ShoppingList{
		UserID: userID,
		Name:   name,
	}
	return uc.repo.CreateList(ctx, list)
}

func (uc *ShoppingListUsecase) GetLists(ctx context.Context, userID uuid.UUID) ([]entity.ShoppingList, error) {
	return uc.repo.GetLists(ctx, userID)
}

func (uc *ShoppingListUsecase) DeleteList(ctx context.Context, listID uuid.UUID, userID uuid.UUID) error {
	return uc.repo.DeleteList(ctx, listID, userID)
}

func (uc *ShoppingListUsecase) UpdateListName(ctx context.Context, listID uuid.UUID, name string) error {
	return uc.repo.UpdateListName(ctx, listID, name)
}

func (uc *ShoppingListUsecase) AddProduct(ctx context.Context, listID uuid.UUID, barcode string) error {
	product := &entity.ShoppingListProduct{
		ShoppingListID: listID,
		Barcode:        barcode,
	}
	return uc.repo.AddProduct(ctx, product)
}

func (uc *ShoppingListUsecase) DeleteProduct(ctx context.Context, productID uuid.UUID, userID uuid.UUID) error {
	return uc.repo.DeleteProduct(ctx, productID, userID)
}

func (uc *ShoppingListUsecase) GetProducts(ctx context.Context, listID uuid.UUID, userID uuid.UUID) ([]entity.ShoppingListProduct, error) {
	return uc.repo.GetProducts(ctx, listID, userID)
}
