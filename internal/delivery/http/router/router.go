package router

import (
	"net/http"

	"SafeBitesServer/internal/delivery/http/handler"
	"SafeBitesServer/internal/delivery/http/middleware"
	"SafeBitesServer/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func NewRouter(
	userUC *usecase.UserUsecase,
	authUC *usecase.AuthUsecase,
	favoritesUC *usecase.FavoritesUsecase,
	listUC *usecase.ShoppingListUsecase) http.Handler {

	r := chi.NewRouter()

	r.Mount("/auth", handler.NewAuthHandler(authUC).Routes())

	userHandler := handler.NewUserHandler(userUC)
	r.Get("/user/all", userHandler.GetAll)

	r.Route("/user", func(r chi.Router) {
		r.Use(middleware.JWTAuth)
		r.Get("/profile", userHandler.GetProfile)
		r.Put("/profile", userHandler.UpdateName)
		r.Delete("/", userHandler.DeleteProfile)
	})

	r.Route("/favorites", func(r chi.Router) {
		r.Use(middleware.JWTAuth)
		r.Mount("/", handler.NewFavoritesHandler(favoritesUC).Routes())
	})

	r.Route("/lists", func(r chi.Router) {
		r.Use(middleware.JWTAuth)
		r.Mount("/", handler.NewShoppingListHandler(listUC).Routes())
	})

	return r
}
