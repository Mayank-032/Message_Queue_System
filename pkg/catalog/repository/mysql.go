package repository

import (
	"context"
	"database/sql"
	"go-message_queue_system/domain/entity"
	"go-message_queue_system/domain/interfaces/repository"
)

type ProductRepo struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) repository.IProductRepo {
	return ProductRepo{
		DB: db,
	}
}

func (pr ProductRepo) Create(ctx context.Context, product entity.Product) (int, error) {
	return 0, nil
}