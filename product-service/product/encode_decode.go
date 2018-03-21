package product

import (
	"context"

	"workshop-go-kit/product-service/proto"
)

func DecodeListRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.ListRequest)
	return ListRequest{req.MaxResult}, nil
}

func EncodeListResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return &proto.ListResponse{
		nil,
	}, nil
}

func DecodeCreateProductRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.CreateProductRequest)
	return CreateProductRequest{req.Name}, nil
}

func EncodeProductResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(ProductResponse)
	return &proto.ProductResponse{
		res.ID,
		res.Name,
	}, nil
}
