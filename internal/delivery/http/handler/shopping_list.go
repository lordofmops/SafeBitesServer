package handler

import (
	"SafeBitesServer/internal/usecase"
	"SafeBitesServer/pkg/response"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ShoppingListHandler struct {
	uc *usecase.ShoppingListUsecase
}

func NewShoppingListHandler(uc *usecase.ShoppingListUsecase) *ShoppingListHandler {
	return &ShoppingListHandler{uc: uc}
}

func (h *ShoppingListHandler) UserRoutes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.CreateList)
	r.Get("/", h.GetLists)
	r.Delete("/{listID}", h.DeleteList)
	r.Put("/{listID}", h.UpdateListName)

	r.Post("/{listID}/products", h.AddProduct)
	r.Get("/{listID}/products", h.GetProducts)
	r.Delete("/products/{productID}", h.DeleteProduct)
	return r
}

func (h *ShoppingListHandler) CreateList(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)
	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "неверный формат запроса")
		return
	}
	if err := h.uc.CreateList(r.Context(), userID, req.Name); err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось создать список")
		return
	}
	response.JSON(w, http.StatusCreated, "список создан")
}

func (h *ShoppingListHandler) GetLists(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)
	lists, err := h.uc.GetLists(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось получить списки")
		return
	}
	response.JSON(w, http.StatusOK, lists)
}

func (h *ShoppingListHandler) DeleteList(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)
	listID, _ := uuid.Parse(chi.URLParam(r, "listID"))
	if err := h.uc.DeleteList(r.Context(), listID, userID); err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось удалить список")
		return
	}
	response.JSON(w, http.StatusOK, "список удален")
}

func (h *ShoppingListHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	listID, _ := uuid.Parse(chi.URLParam(r, "listID"))
	var req struct {
		Barcode string `json:"barcode"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "неверный формат запроса")
		return
	}
	if err := h.uc.AddProduct(r.Context(), listID, req.Barcode); err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось добавить продукт")
		return
	}
	response.JSON(w, http.StatusCreated, "продукт добавлен")
}

func (h *ShoppingListHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)

	listID, err := uuid.Parse(chi.URLParam(r, "listID"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "невалидный listID")
		return
	}

	products, err := h.uc.GetProducts(r.Context(), listID, userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось получить продукты")
		return
	}
	response.JSON(w, http.StatusOK, products)
}

func (h *ShoppingListHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)
	productID, _ := uuid.Parse(chi.URLParam(r, "productID"))
	if err := h.uc.DeleteProduct(r.Context(), productID, userID); err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось удалить продукт")
		return
	}
	response.JSON(w, http.StatusOK, "продукт удален")
}

func (h *ShoppingListHandler) UpdateListName(w http.ResponseWriter, r *http.Request) {
	listIDStr := chi.URLParam(r, "listID")
	listID, err := uuid.Parse(listIDStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "невалидный ID списка")
		return
	}

	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "невалидный формат запроса")
		return
	}

	if err := h.uc.UpdateListName(r.Context(), listID, req.Name); err != nil {
		response.Error(w, http.StatusInternalServerError, "не удалось обновить имя списка")
		return
	}

	response.JSON(w, http.StatusOK, "имя списка обновлено")
}
