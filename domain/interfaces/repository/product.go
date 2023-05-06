package repository

import (
	"context"
	"go-message_queue_system/domain/entity"
)

type IProductRepo interface {
	Create(ctx context.Context, product entity.Product) (int, error)
}