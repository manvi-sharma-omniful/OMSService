package services

import (
	"context"
	"log"

	"github.com/omniful/go_commons/kafka"
	"github.com/omniful/go_commons/pubsub"
)

type KafkaProducer struct {
	Producer *kafka.ProducerClient
}

var producerKafka *KafkaProducer

func SetProducer(producer *kafka.ProducerClient) {
	if producerKafka == nil {
		producerKafka = &KafkaProducer{}
	}
	producerKafka.Producer = producer
}

func getProducer() *kafka.ProducerClient {
	if producerKafka != nil {
		return producerKafka.Producer
	}
	return nil
}

func PublishMessageToKafka(bytesOrderItem []byte, orderID string) {
	ctx := context.Background()
	msg := &pubsub.Message{
		Topic: "oms-service-topic",
		Key:   orderID,
		Value: bytesOrderItem,
		Headers: map[string]string{
			"custom-header": "value",
		},
	}
	producer := getProducer()
	err := producer.Publish(ctx, msg)
	if err != nil {
		log.Println("Error publishing message to kafka")
	} else {
		log.Println("Message published to kafka")
	}
}
