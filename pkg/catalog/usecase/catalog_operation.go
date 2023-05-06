package usecase

import (
	"context"
	"go-message_queue_system/domain/entity"
	"go-message_queue_system/domain/interfaces/repository"
	"go-message_queue_system/domain/interfaces/usecase"
)

type ProductUCase struct {
	ProductRepo repository.IProductRepo
}

func NewProductUCase(productRepo repository.IProductRepo) usecase.IProductUCase {
	return ProductUCase {
		ProductRepo: productRepo,
	}
}

func (puc ProductUCase) CreateProduct(ctx context.Context, product entity.Product) (int, error) {
	return 0, nil
}