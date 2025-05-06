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
	listUC *usecase.ShoppingListUsecase,
	storeUC *usecase.StoreUsecase,
	historyUC *usecase.SearchHistoryUsecase,
	restrictionUC *usecase.RestrictionUsecase) http.Handler {

	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Mount("/auth", handler.NewAuthHandler(authUC).Routes())

		r.Route("/public", func(r chi.Router) {
			r.Mount("/stores", handler.NewStoreHandler(storeUC).PublicRoutes())
			r.Mount("/restrictions", handler.NewRestrictionHandler(restrictionUC).PublicRoutes())
			r.Mount("/user", handler.NewUserHandler(userUC).PublicRoutes())
		})

		r.Route("/me", func(r chi.Router) {
			r.Use(middleware.JWTAuth)

			r.Mount("/user", handler.NewUserHandler(userUC).UserRoutes())
			r.Mount("/favorites", handler.NewFavoritesHandler(favoritesUC).UserRoutes())
			r.Mount("/history", handler.NewSearchHistoryHandler(historyUC).UserRoutes())
			r.Mount("/lists", handler.NewShoppingListHandler(listUC).UserRoutes())
			r.Mount("/restrictions", handler.NewRestrictionHandler(restrictionUC).UserRoutes())
		})
	})

	return r
}
