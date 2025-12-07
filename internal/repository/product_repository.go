package repository

import (
	"task-api/internal/database"
	"task-api/internal/models"
)

type ProductRepo struct{}

var Product = ProductRepo{}

func (ProductRepo) GetAll() ([]models.Product, error) {
	var Product []models.Product
	result := database.DB.Find(&Product)
	return Product, result.Error
}

func (ProductRepo) Create(product *models.Product) error {
	return database.DB.Create(product).Error
}
