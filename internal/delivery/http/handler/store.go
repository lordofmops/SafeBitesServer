package handler

import (
	"encoding/json"
	"net/http"

	"SafeBitesServer/internal/usecase"
	"SafeBitesServer/pkg/response"
	"github.com/go-chi/chi/v5"
)

type StoreHandler struct {
	uc *usecase.StoreUsecase
}

func NewStoreHandler(uc *usecase.StoreUsecase) *StoreHandler {
	return &StoreHandler{uc: uc}
}

func (h *StoreHandler) PublicRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/all", h.GetAll)
	r.Post("/", h.CreateStore)
	return r
}

func (h *StoreHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	stores, err := h.uc.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось получить магазины")
		return
	}
	response.JSON(w, http.StatusOK, stores)
}

func (h *StoreHandler) CreateStore(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
		Link string `json:"link"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "неверный формат запроса")
		return
	}
	if err := h.uc.CreateStore(r.Context(), req.Name, req.Link); err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось добавить магазин")
		return
	}
	response.JSON(w, http.StatusCreated, "магазин добавлен")
}
