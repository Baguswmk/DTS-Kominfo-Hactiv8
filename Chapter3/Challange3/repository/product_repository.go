package repository

import "DTS-Kominfo-Hactiv8/Chapter3/Challange3/models"

//go:generate mockery --name entity.Product
type ProductRepository interface {
	FindById(id string) *models.Product
	FindAll() []*models.Product
}
