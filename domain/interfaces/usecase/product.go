package usecase

import (
	"context"
	"go-message_queue_system/domain/entity"
)

type IProductUCase interface {
	UpsertProduct(ctx context.Context, product entity.Product) error
}