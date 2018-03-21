package product

import "workshop-go-kit/product-service/database"

type Service interface {
	GetList() ([]*database.Product, error)
	CreateProduct(name string) (*database.Product, error)
}

type ProductService struct {
	Repository database.ProductRepository
}

func (s *ProductService) GetList() ([]*database.Product, error) {
	return s.Repository.FindAll()
}

func (s *ProductService) CreateProduct(name string) (*database.Product, error) {
	return s.Repository.New(name)
}
