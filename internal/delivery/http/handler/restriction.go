package handler

import (
	"SafeBitesServer/internal/usecase"
	"SafeBitesServer/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"

	"encoding/json"
)

type RestrictionHandler struct {
	uc *usecase.RestrictionUsecase
}

func NewRestrictionHandler(uc *usecase.RestrictionUsecase) *RestrictionHandler {
	return &RestrictionHandler{uc: uc}
}

func (h *RestrictionHandler) UserRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.GetUserRestrictions)
	r.Post("/", h.Add)
	r.Delete("/{id}", h.Remove)
	return r
}

func (h *RestrictionHandler) PublicRoutes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.CreateRestriction)
	r.Get("/all", h.GetAll)
	return r
}

func (h *RestrictionHandler) GetUserRestrictions(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)
	restrictions, err := h.uc.GetUserRestrictions(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось получить ограничения")
		return
	}
	response.JSON(w, http.StatusOK, restrictions)
}

func (h *RestrictionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	restrictions, err := h.uc.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось получить ограничения")
		return
	}
	response.JSON(w, http.StatusOK, restrictions)
}

func (h *RestrictionHandler) CreateRestriction(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
		Type string `json:"type"`
		Tag  string `json:"tag"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "неверный формат запроса")
		return
	}
	if err := h.uc.CreateRestriction(r.Context(), req.Name, req.Type, req.Tag); err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось добавить ограничение")
		return
	}
	response.JSON(w, http.StatusCreated, "ограничение добавлено")
}

func (h *RestrictionHandler) Add(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)

	var req struct {
		RestrictionID uuid.UUID `json:"restriction_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "неверный формат запроса")
		return
	}

	restrictions, err := h.uc.Add(r.Context(), userID, req.RestrictionID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось добавить")
		return
	}
	response.JSON(w, http.StatusCreated, restrictions)
}

func (h *RestrictionHandler) Remove(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)
	restrictionID, _ := uuid.Parse(chi.URLParam(r, "id"))

	restrictions, err := h.uc.Remove(r.Context(), userID, restrictionID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось удалить")
		return
	}
	response.JSON(w, http.StatusOK, restrictions)
}
