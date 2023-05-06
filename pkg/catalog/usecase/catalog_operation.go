package usecase

import (
	"context"
	"errors"
	"go-message_queue_system/domain/entity"
	"go-message_queue_system/domain/interfaces/repository"
	"go-message_queue_system/domain/interfaces/usecase"
	"log"
)

type ProductUCase struct {
	ProductRepo repository.IProductRepo
}

func NewProductUCase(productRepo repository.IProductRepo) usecase.IProductUCase {
	return ProductUCase {
		ProductRepo: productRepo,
	}
}

func (puc ProductUCase) UpsertProduct(ctx context.Context, product entity.Product) error {
	productId, err := puc.ProductRepo.Upsert(ctx, product)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_upsert_product_in_database\n\n", err.Error())
		return errors.New("unable to upsert product")
	}
	err = publishProductIdToQueue(ctx, productId)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_publish_data_to_queue\n\n", err.Error())
		return errors.New("unable to publish to queue")
	}
	return nil
}
