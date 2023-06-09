package usecase

import (
	"context"
	"errors"
	"go-message_queue_system/domain/entity"
	"go-message_queue_system/domain/interfaces/repository"
	"go-message_queue_system/domain/interfaces/usecase"
	"log"

	"github.com/streadway/amqp"
)

var (
	publishProductIdToQueue = PublishProductIdToQueue
)

type ProductUCase struct {
	Conn        *amqp.Connection
	UserUCase   usecase.IUserUCase
	ProductRepo repository.IProductRepo
}

func NewProductUCase(conn *amqp.Connection, userUCase usecase.IUserUCase, productRepo repository.IProductRepo) usecase.IProductUCase {
	return ProductUCase{
		Conn:        conn,
		UserUCase:   userUCase,
		ProductRepo: productRepo,
	}
}

func (puc ProductUCase) UpsertProduct(ctx context.Context, product entity.Product) error {
	userExists, err := puc.UserUCase.FetchUser(ctx, product.UserId)
	if err != nil || !userExists {
		log.Printf("Error: %v\n, \n\n", err.Error())
		return err
	}

	productId, err := puc.ProductRepo.Upsert(ctx, product)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_upsert_product_in_database\n\n", err.Error())
		return errors.New("unable to upsert product")
	}
	err = publishProductIdToQueue(ctx, puc.Conn, productId)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_publish_data_to_queue\n\n", err.Error())
		return errors.New("unable to publish to queue")
	}
	return nil
}
