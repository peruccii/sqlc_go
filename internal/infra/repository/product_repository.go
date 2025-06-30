package repository

import (
	"context"
	"database/sql"

	"golearn/db"
)

type ProductRepository struct {
	queries *db.Queries
}

// like constructor
func NewProductRepositoryMysql(dbConn *sql.DB) *ProductRepository {
	return &ProductRepository{
		queries: db.New(dbConn),
	}
}

// creating a method and it receiver will be ProductRepository
func (r *ProductRepository) CreateProduct(ctx context.Context, input db.CreateProductParams) error {
	return r.queries.CreateProduct(ctx, input)
}

func (r *ProductRepository) ListProduct(ctx context.Context) ([]db.Product, error) {
	return r.queries.ListProducts(ctx)
}
