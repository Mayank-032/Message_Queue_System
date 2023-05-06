package usecase

import (
	"context"
	"go-message_queue_system/domain/entity"
)

type IProductUCase interface {
	CreateProduct(ctx context.Context, product entity.Product) (int, error)
}