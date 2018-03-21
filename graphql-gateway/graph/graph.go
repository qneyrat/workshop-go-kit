package graph

import (
	"context"
	"fmt"
	"math/rand"
)

type App struct {
	products []Product
}

func (a *App) Query_products(ctx context.Context) ([]Product, error) {
	return a.products, nil
}

func (a *App) Mutation_createProduct(ctx context.Context, name string) (Product, error) {
	product := Product{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Name: name,
	}
	a.products = append(a.products, product)
	return product, nil
}
