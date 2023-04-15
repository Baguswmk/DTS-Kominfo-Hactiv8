package service

import (
	"DTS-Kominfo-Hactiv8/Chapter3/Challange3/models"
	"DTS-Kominfo-Hactiv8/Chapter3/Challange3/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductService_GetAllProduct(t *testing.T) {

	productRepository := mocks.NewProductRepository(t)

	os := &ProductService{
		ProductRepository: productRepository,
	}

	mockData := []*models.Product{
		{
			Title:   "1",
			Description: "Cheese",
		},
	}

	productRepository.On("FindAll").Return(mockData, nil)

	got, _ := os.GetAllProduct()

	assert.Equal(t, mockData, got, "Error response must be nil")
}

func TestProductService_GetAllProductNotFound(t *testing.T) {

	productRepository := mocks.NewProductRepository(t)

	os := &ProductService{
		ProductRepository: productRepository,
	}

	productRepository.On("FindAll").Return(nil, nil)

	got, err := os.GetAllProduct()
	assert.Nil(t, got)
	assert.NotNil(t, err)
	assert.Equal(t, "all product not found", err.Error(), "error response has to be 'all product not found'")
}

func TestProductService_GetProductById(t *testing.T) {

	productRepository := mocks.NewProductRepository(t)

	mockData := &models.Product{
		Title:   "1",
		Description: "Cheese",
	}
	productRepository.On("FindById", "1").Return(mockData, nil)

	os := &ProductService{
		ProductRepository: productRepository,
	}

	got, _ := os.GetOneProductId("1")
	assert.Equal(t, mockData, got, "error response must be nil")
}

func TestProductService_GetProductByIdNotFound(t *testing.T) {

	productRepository := mocks.NewProductRepository(t)

	productRepository.On("FindById", "1").Return(nil, nil)

	os := &ProductService{
		ProductRepository: productRepository,
	}

	got, err := os.GetOneProductId("1")
	assert.Nil(t, got)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}
