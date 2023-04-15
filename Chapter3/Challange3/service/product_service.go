package service

import (
	"errors"
	"DTS-Kominfo-Hactiv8/Chapter3/Challange3/models"
	"DTS-Kominfo-Hactiv8/Chapter3/Challange3/repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (service ProductService) GetOneProductId(id string) (*models.Product, error) {
	product := service.ProductRepository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (service ProductService) GetAllProduct() ([]*models.Product, error) {
	product := service.ProductRepository.FindAll()
	if product == nil {
		return nil, errors.New("all product not found")
	}

	return product, nil
}
