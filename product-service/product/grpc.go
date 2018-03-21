package product

import (
	"context"

	transport "github.com/go-kit/kit/transport/grpc"

	"workshop-go-kit/product-service/proto"
)

type Server struct {
	list          transport.Handler
	createProduct transport.Handler
}

func (s *Server) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	_, resp, err := s.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.ListResponse), nil
}

func (s *Server) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.ProductResponse, error) {
	_, resp, err := s.createProduct.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.ProductResponse), nil
}

func NewServer(_ context.Context, endpoint Endpoints) proto.ProductServer {
	return &Server{
		list: transport.NewServer(
			endpoint.ListEndpoint,
			DecodeListRequest,
			EncodeListResponse,
		),
		createProduct: transport.NewServer(
			endpoint.CreateProductEndpoint,
			DecodeCreateProductRequest,
			EncodeProductResponse,
		),
	}
}
