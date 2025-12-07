package service

import (
	"task-api/internal/models"
	"task-api/internal/repository"
)

type ProductService struct{}

var Product = ProductService{}

func (ProductService) GetAll() ([]models.Product, error) {
	return repository.Product.GetAll()
}

func (ProductService) Create(name string, description string, price uint) (models.Product, error) {
	product := models.Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	err := repository.Product.Create(&product)
	return product, err
}
