package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"workshop-go-kit/product-service/database"
	"workshop-go-kit/product-service/product"
	"workshop-go-kit/product-service/proto"
)

func main() {
	productService := &product.ProductService{database.PGProductRepository{}}
	endpoints := product.Endpoints{
		ListEndpoint:          product.MakeListEndpoint(productService),
		CreateProductEndpoint: product.MakeCreateProductEndpoint(productService),
	}

	ctx := context.Background()
	handler := product.NewServer(ctx, endpoints)
	server := grpc.NewServer()
	proto.RegisterProductServer(server, handler)

	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("gRPC listen on :4000")
	server.Serve(listener)
}
