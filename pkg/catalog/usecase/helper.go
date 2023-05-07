package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"go-message_queue_system/rabbitmq/publisher"
	"log"

	"github.com/streadway/amqp"
)

func publishProductIdToQueue(ctx context.Context, conn *amqp.Connection, productId int) error {
	defer conn.Close()
	amqpChannel, err := conn.Channel()
	if err != nil {
		log.Printf("Error: %v,\n failed_to_create_channel", err.Error())
		return errors.New("unable to create channel")
	}
	defer amqpChannel.Close()

	publishData := publisher.PublishTaskRequestData{}
	publishData.Data = productId
	reqBytes, err := json.Marshal(publishData)
	if err != nil {
		log.Printf("Error: %v,\n invalid_json_format", err.Error())
		return errors.New("invalid json format")
	}

	publishRequest := publisher.PublishTaskRequest{}
	publishRequest.QueueName = "store_product_images"
	publishRequest.ExchangeName = "store_product"
	publishRequest.RoutingKey = "store_product_images"
	publishRequest.ReqBytes = reqBytes
	err = publishRequest.PublishMessage(ctx, amqpChannel)
	if err != nil {
		log.Printf("Error: %v,\n failed_to_publish_message\n\n", err.Error())
		return errors.New("unable to publish message")
	}
	return nil
}