package service

import (
	"context"
	"fmt"
	"log"

	"awesomeProject/Project/OMS/domain"
	"awesomeProject/Project/OMS/repository"
	"github.com/omniful/go_commons/sqs"
	"go.mongodb.org/mongo-driver/mongo"
)

var Collection = &repository.OrderCollection{}

func SetupOrderCollection(collection *mongo.Collection) {
	Collection.Collection = collection
}

var Producer = &sqs.Publisher{}

func SetProducer(ctx context.Context, q *sqs.Queue) {
	Producer = sqs.NewPublisher(q)
}

func createMessage(path string) *sqs.Message {
	return &sqs.Message{
		GroupId:       "1",
		Value:         []byte(path),
		ReceiptHandle: "orders",
	}
}

func ConvertControllerRequestToService(ctx context.Context, path string) error {
	message := createMessage(path)
	err := SendMessage(ctx, message)
	if err != nil {
		return err
	}
	return nil
}

func SendMessage(ctx context.Context, message *sqs.Message) error {
	err := Producer.Publish(ctx, message)
	if err != nil {
		log.Fatal("Not received", err)
		return err
	}
	fmt.Println("Message Published")
	return nil
}

func GetOrders(filePath string) {
	orders, err := CSVOperation(filePath)
	if err != nil {
		fmt.Println("\nfailed to parse csv with path : ", filePath)
		return
	}
	for _, order := range orders {
		PlaceOrder(*order)
	}

}

func PlaceOrder(order domain.Order) {
	Collection.Create(order)
}
