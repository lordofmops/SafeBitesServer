package handler

import (
	"SafeBitesServer/internal/usecase"
	"SafeBitesServer/pkg/response"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type FavoritesHandler struct {
	uc *usecase.FavoritesUsecase
}

func NewFavoritesHandler(uc *usecase.FavoritesUsecase) *FavoritesHandler {
	return &FavoritesHandler{uc: uc}
}

func (h *FavoritesHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.List)
	r.Post("/", h.Add)
	r.Delete("/", h.Delete)
	return r
}

func (h *FavoritesHandler) Add(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)

	var req struct {
		Barcode string `json:"barcode"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "неверный формат данных")
		return
	}

	if err := h.uc.Add(r.Context(), userID, req.Barcode); err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось добавить")
		return
	}
	response.JSON(w, http.StatusCreated, "добавлено")
}

func (h *FavoritesHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)

	var req struct {
		Barcode string `json:"barcode"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "неверный формат данных")
		return
	}

	if err := h.uc.Delete(r.Context(), userID, req.Barcode); err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось удалить")
		return
	}
	response.JSON(w, http.StatusOK, "удалено")
}

func (h *FavoritesHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)

	barcodes, err := h.uc.List(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "ошибка получения")
		return
	}
	response.JSON(w, http.StatusOK, barcodes)
}
