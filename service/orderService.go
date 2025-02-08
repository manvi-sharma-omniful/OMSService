package service

import (
	"awesomeProject/Project/OMS/domain"
	"awesomeProject/Project/OMS/repository"
	"context"
	"fmt"
	"github.com/omniful/go_commons/sqs"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var Collection = &repository.OrderCollection{}

func SetupOrderCollection(collection *mongo.Collection) {
	Collection.Collection = collection
}

var Producer = &sqs.Publisher{}

func SetProducer(ctx context.Context, q *sqs.Queue) {
	// Producer = p
	Producer = sqs.NewPublisher(q)
}

func createMessage(path string) *sqs.Message {
	return &sqs.Message{
		GroupId:       "group-123",
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
		log.Fatal("ni aaya", err)
		return err
	}
	fmt.Println("Message Published")
	return nil
}

func GetOrders(filePath string) {
	orders, err := CSVOperation(filePath)
	if err != nil {
		// c.JSON(500, gin.H{"error": err.Error()})
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
