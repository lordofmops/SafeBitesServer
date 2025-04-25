package repository

import (
	"SafeBitesServer/internal/entity"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShoppingListRepository struct {
	db *gorm.DB
}

func NewShoppingListRepository(db *gorm.DB) *ShoppingListRepository {
	return &ShoppingListRepository{db: db}
}

func (r *ShoppingListRepository) CreateList(ctx context.Context, list *entity.ShoppingList) error {
	return r.db.WithContext(ctx).Create(list).Error
}

func (r *ShoppingListRepository) GetLists(ctx context.Context, userID uuid.UUID) ([]entity.ShoppingList, error) {
	var lists []entity.ShoppingList
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&lists).Error
	return lists, err
}

func (r *ShoppingListRepository) DeleteList(ctx context.Context, listID uuid.UUID, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Where("id = ? AND user_id = ?", listID, userID).Delete(&entity.ShoppingList{}).Error
}

func (r *ShoppingListRepository) UpdateListName(ctx context.Context, listID uuid.UUID, name string) error {
	return r.db.WithContext(ctx).Model(&entity.ShoppingList{}).
		Where("id = ?", listID).
		Update("name", name).Error
}

func (r *ShoppingListRepository) AddProduct(ctx context.Context, product *entity.ShoppingListProduct) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *ShoppingListRepository) DeleteProduct(ctx context.Context, productID uuid.UUID, userID uuid.UUID) error {
	subQuery := r.db.Model(&entity.ShoppingList{}).Select("id").Where("user_id = ?", userID)
	return r.db.WithContext(ctx).
		Where("id = ? AND shopping_list_id IN (?)", productID, subQuery).
		Delete(&entity.ShoppingListProduct{}).Error
}

func (r *ShoppingListRepository) GetProducts(ctx context.Context, listID uuid.UUID, userID uuid.UUID) ([]entity.ShoppingListProduct, error) {
	var products []entity.ShoppingListProduct
	err := r.db.WithContext(ctx).
		Joins("JOIN shopping_lists ON shopping_lists.id = shopping_list_products.shopping_list_id").
		Where("shopping_lists.user_id = ? AND shopping_lists.id = ?", userID, listID).
		Find(&products).Error
	return products, err
}
