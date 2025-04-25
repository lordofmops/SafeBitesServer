package main

import (
	"SafeBitesServer/internal/entity"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"SafeBitesServer/internal/delivery/http/router"
	"SafeBitesServer/internal/repository"
	"SafeBitesServer/internal/usecase"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=safebites port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	db.AutoMigrate(
		&entity.User{},
		&entity.Favorites{},
		&entity.ShoppingList{},
		&entity.ShoppingListProduct{},
		&entity.Store{})

	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)

	favoritesRepo := repository.NewFavoriteRepository(db)
	favoritesUC := usecase.NewFavoritesUsecase(favoritesRepo)

	authUC := usecase.NewAuthUsecase(userRepo, []byte("your-secret-key"))

	listRepo := repository.NewShoppingListRepository(db)
	listUC := usecase.NewShoppingListUsecase(listRepo)

	storeRepo := repository.NewStoreRepository(db)
	storeUC := usecase.NewStoreUsecase(storeRepo)

	r := router.NewRouter(userUC, authUC, favoritesUC, listUC, storeUC)

	log.Println("Server starting at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
