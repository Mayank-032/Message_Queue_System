package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-message_queue_system/domain/entity"
	"go-message_queue_system/domain/interfaces/repository"
	"log"
	"encoding/json"
)

type ProductRepo struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) repository.IProductRepo {
	return ProductRepo{
		DB: db,
	}
}

func (pr ProductRepo) Upsert(ctx context.Context, product entity.Product) (int, error) {
	conn, err := pr.DB.Conn(ctx)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_db_connect\n\n", err.Error())
		return 0, errors.New("unable to db connect")
	}
	defer conn.Close()

	sqlQuery := "INSERT INTO product(product_name, product_description, product_price, product_images)" + 
	" VALUES(?, ?, ?, ?) ON DUPLICATE KEY UPDATE" + 
	" product_description=values(product_description), product_price=values(product_price), product_images=values(product_images)"

	productImageBytes, err := json.Marshal(product.Images)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_marshal_array_to_json\n\n", err.Error())
		return 0, errors.New("unable to marshal array to json")
	}

	args := []interface{}{product.Name, product.Description, product.Price, productImageBytes}
	result, err := conn.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_execute_sql_query\n\n", err.Error())
		return 0, errors.New("unable to execute sql query")
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error: %v\n, unable_to_fetch_last_insertedId\n\n", err.Error())
		return 0, errors.New("unable to fetch last insertedId")
	}
	return int(lastInsertedId), nil
}