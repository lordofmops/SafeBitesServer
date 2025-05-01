package handler

import (
	"SafeBitesServer/internal/usecase"
	"SafeBitesServer/pkg/response"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type SearchHistoryHandler struct {
	uc *usecase.SearchHistoryUsecase
}

func NewSearchHistoryHandler(uc *usecase.SearchHistoryUsecase) *SearchHistoryHandler {
	return &SearchHistoryHandler{uc: uc}
}

func (h *SearchHistoryHandler) UserRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.GetAll)
	r.Post("/", h.Add)
	r.Delete("/", h.Clear)
	r.Delete("/{id}", h.Delete)
	return r
}

func (h *SearchHistoryHandler) Add(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)

	var req struct {
		Barcode string `json:"barcode"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Barcode == "" {
		response.Error(w, http.StatusBadRequest, "неверный формат запроса")
		return
	}

	if err := h.uc.Add(r.Context(), userID, req.Barcode); err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось сохранить запрос")
		return
	}

	response.JSON(w, http.StatusCreated, "запрос добавлен в историю")
}

func (h *SearchHistoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)

	history, err := h.uc.GetAll(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось получить историю")
		return
	}

	response.JSON(w, http.StatusOK, history)
}

func (h *SearchHistoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "невалидный ID")
		return
	}

	if err := h.uc.Delete(r.Context(), id, userID); err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось удалить запись")
		return
	}

	response.JSON(w, http.StatusOK, "запись удалена")
}

func (h *SearchHistoryHandler) Clear(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)

	if err := h.uc.Clear(r.Context(), userID); err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось очистить историю")
		return
	}

	response.JSON(w, http.StatusOK, "история очищена")
}
