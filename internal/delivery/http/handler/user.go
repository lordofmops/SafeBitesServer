package handler

import (
	"encoding/json"
	"net/http"

	"SafeBitesServer/internal/usecase"
	"SafeBitesServer/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserHandler struct {
	uc *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) UserRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/profile", h.GetProfile)
	r.Put("/profile", h.UpdateName)
	r.Delete("/", h.DeleteProfile)
	return r
}

func (h *UserHandler) PublicRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/all", h.GetAll)
	return r
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.uc.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось получить пользователей")
		return
	}
	response.JSON(w, http.StatusOK, users)
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)

	user, err := h.uc.GetProfile(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось получить профиль")
		return
	}
	response.JSON(w, http.StatusOK, user)
}

func (h *UserHandler) UpdateName(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)

	var data struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		response.Error(w, http.StatusBadRequest, "неправильный формат данных")
		return
	}

	user, err := h.uc.UpdateName(r.Context(), userID, data.Name)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось обновить имя")
		return
	}
	response.JSON(w, http.StatusOK, user)
}

func (h *UserHandler) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)

	if err := h.uc.DeleteProfile(r.Context(), userID); err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось удалить профиль")
		return
	}
	response.JSON(w, http.StatusOK, "профиль удален")
}
