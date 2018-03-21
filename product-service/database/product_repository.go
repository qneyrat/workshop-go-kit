package database

type Product struct {
	ID   int32
	Name string
}

type ProductRepository interface {
	FindAll() ([]*Product, error)
	New(name string) (*Product, error)
}

type PGProductRepository struct{}

func (r PGProductRepository) FindAll() ([]*Product, error) {
	product1 := &Product{1, "test"}
	product2 := &Product{2, "test"}

	return []*Product{product1, product2}, nil
}

func (r PGProductRepository) New(name string) (*Product, error) {
	return &Product{1, name}, nil
}
