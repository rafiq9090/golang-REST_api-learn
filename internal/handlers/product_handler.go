package handlers

import (
	"encoding/json"
	"net/http"
	"task-api/internal/dto"
	"task-api/internal/service"
	"task-api/internal/utils"
)

// GetProducts godoc
// @Summary      Get all products (protected)
// @Tags         products
// @Security     BearerAuth
// @Produce      json
// @Success      200 {array} models.Product
// @Failure      401 {object} utils.ErrorResponse
// @Router       /products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := service.Product.GetAll()
	if err != nil {
		utils.JSONError(w, "Failed to fetch products", http.StatusInternalServerError, nil)
		return
	}
	utils.JSONSuccess(w, products)
}

// CreateProduct godoc
// @Summary      Create a new product (protected)
// @Tags         products
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body dto.CreateProductRequest true "Product data"
// @Success      201 {object} models.Product
// @Failure      400,401 {object} utils.ErrorResponse
// @Router       /products [post]
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.JSONError(w, "Invalid Request Body: "+err.Error(), http.StatusBadRequest, nil)
		return
	}
	if errs := dto.ValidateCreateProduct(input); errs != nil {
		utils.JSONError(w, "Validation Failed", http.StatusBadRequest, errs)
		return
	}

	product, err := service.Product.Create(input.Name, input.Description, input.Price)
	if err != nil {
		utils.JSONError(w, "Failed to create product", http.StatusInternalServerError, nil)
		return
	}
	utils.JSONSuccess(w, product)
}
