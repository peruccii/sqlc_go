package repository

import (
	"context"
	"database/sql"

	"golearn/db"

	"github.com/google/uuid"
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

func (r *ProductRepository) GetProduct(ctx context.Context, id uuid.UUID) (db.Product, error) {
	return r.queries.GetProduct(ctx, id)
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	return r.queries.DeleteProduct(ctx, id)
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, input db.UpdateProductParams) error {
	return r.queries.UpdateProduct(ctx, input)
}
