package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type ProductResponse struct {
	ID   int32
	Name string
}

type ListRequest struct {
	MaxResult int32
}

type ListReponse struct {
	ProductList []*ProductResponse
}

type CreateProductRequest struct {
	Name string
}

type Endpoints struct {
	ListEndpoint          endpoint.Endpoint
	CreateProductEndpoint endpoint.Endpoint
}

func MakeListEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		list, _ := svc.GetList()
		productList := []*ProductResponse{}
		for index := range list {
			product := list[index]
			productList[index] = &ProductResponse{
				product.ID,
				product.Name,
			}
		}

		return ListReponse{ProductList: productList}, nil
	}
}

func MakeCreateProductEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		cr := req.(CreateProductRequest)
		product, _ := svc.CreateProduct(cr.Name)

		return ProductResponse{
			product.ID,
			product.Name,
		}, nil
	}
}
