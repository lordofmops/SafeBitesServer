package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"

	"SafeBitesServer/internal/usecase"
	"SafeBitesServer/pkg/response"
)

type AuthHandler struct {
	uc *usecase.AuthUsecase
}

func NewAuthHandler(uc *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

func (h *AuthHandler) Routes() http.Handler {
	mux := chi.NewRouter()
	mux.Post("/register", h.Register)
	mux.Post("/login", h.Login)
	return mux
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "неверный формат")
		return
	}
	user, err := h.uc.Register(r.Context(), req.Login, req.Password, req.Name)
	if err != nil {
		response.Error(w, http.StatusConflict, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, user)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "неверный формат")
		return
	}
	token, err := h.uc.Login(r.Context(), req.Login, req.Password)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"token": token})
}

//func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
//	userIDRaw := r.Context().Value("userID")
//	userID, ok := userIDRaw.(uuid.UUID)
//	if !ok {
//		response.Error(w, http.StatusUnauthorized, "некорректный токен")
//		return
//	}
//
//	user, err := h.uc.Me(r.Context(), userID)
//	if err != nil {
//		response.Error(w, http.StatusNotFound, "пользователь не найден")
//		return
//	}
//
//	response.JSON(w, http.StatusOK, user)
//}
