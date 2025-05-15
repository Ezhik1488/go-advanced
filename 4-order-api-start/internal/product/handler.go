package product

import (
	"gorm.io/gorm"
	"net/http"
	"order-api/pkg/middleware"
	"order-api/pkg/req"
	"order-api/pkg/res"
	"strconv"
)

type Response map[string]interface{}

type ProductHandler struct {
	ProductRepo *ProductRepository
}

type ProductHandlerDeps struct {
	ProductRepo *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps *ProductHandlerDeps) *ProductHandler {
	handler := &ProductHandler{
		ProductRepo: deps.ProductRepo,
	}
	router.Handle("GET /product", middleware.Auth(handler.GetALL()))
	router.HandleFunc("GET /product/{id}", handler.GetByID())
	router.Handle("POST /product", middleware.Auth(handler.Create()))
	router.Handle("PATCH /product/{id}", middleware.Auth(handler.Update()))
	router.Handle("DELETE /product/{id}", middleware.Auth(handler.Delete()))
	return handler
}

func (h *ProductHandler) GetALL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := h.ProductRepo.FindALL()
		if err != nil {
			res.JSON(w, Response{
				"result": err.Error(),
				"status": http.StatusNotFound},
				http.StatusNotFound)
			return
		}
		res.JSON(w, Response{
			"result": result,
			"status": http.StatusOK},
			http.StatusOK)
	}
}

func (h *ProductHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			res.JSON(w, Response{
				"result": err.Error(),
				"status": http.StatusBadRequest},
				http.StatusBadRequest)
			return
		}
		product, err := h.ProductRepo.FindByID(uint(id))
		if err != nil {
			res.JSON(w, Response{
				"result": err.Error(),
				"status": http.StatusNotFound},
				http.StatusNotFound)
			return
		}
		res.JSON(w, Response{
			"result": product,
			"status": http.StatusOK},
			http.StatusOK)
	}
}

func (h *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductCreateRequest](w, r)
		if err != nil {
			return
		}
		product := NewProduct(body)
		err = h.ProductRepo.Create(product)
		if err != nil {
			res.JSON(w, Response{
				"result": err.Error(),
				"status": http.StatusBadRequest},
				http.StatusBadRequest)
			return
		}
		res.JSON(w, Response{
			"result": product,
			"status": http.StatusOK},
			http.StatusOK)
	}
}

func (h *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			res.JSON(w, Response{
				"result": err.Error(),
				"status": http.StatusBadRequest},
				http.StatusBadRequest)
			return
		}
		body, err := req.HandleBody[ProductUpdateRequest](w, r)
		if err != nil {
			return
		}
		result, err := h.ProductRepo.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        body.Name,
			Price:       body.Price,
			Description: body.Description,
			Image:       body.Image,
		})
		if err != nil {
			res.JSON(w, Response{
				"result": err.Error(),
				"status": http.StatusInternalServerError},
				http.StatusInternalServerError)
		}
		res.JSON(w, Response{
			"result": result,
			"status": http.StatusOK},
			http.StatusOK)
	}
}

func (h *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			res.JSON(w, Response{
				"result": err.Error(),
				"status": http.StatusBadRequest},
				http.StatusBadRequest)
			return
		}
		result, err := h.ProductRepo.Delete(uint(id))
		if err != nil {
			res.JSON(w, Response{
				"result": err.Error(),
				"status": http.StatusInternalServerError},
				http.StatusInternalServerError)
			return
		}
		if result == 0 {
			res.JSON(w, Response{
				"result": "row not found",
				"status": http.StatusNotFound},
				http.StatusNotFound)
			return
		}
		res.JSON(w, Response{
			"result": "success",
			"status": http.StatusNoContent},
			http.StatusNoContent)
	}
}
