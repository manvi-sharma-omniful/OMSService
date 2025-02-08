package init

import (
	"awesomeProject/Project/OMS/pkg/db"
	"awesomeProject/Project/OMS/service"
	"context"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/kafka"
	"github.com/omniful/go_commons/log"
)

func Initialize(ctx context.Context) {
	initializeLog(ctx)
	initializeDB(ctx)
	initializeSQSProducer(ctx)
	initializeSQSConsumer(ctx)
	initializeKafkaProducer(ctx)

}

func initializeSQSConsumer(ctx context.Context) {

}

func initializeSQSProducer(ctx context.Context) {

}

// Initialize logging
func initializeLog(ctx context.Context) {
	err := log.InitializeLogger(
		log.Formatter(config.GetString(ctx, "log.format")),
		log.Level(config.GetString(ctx, "log.level")),
	)
	if err != nil {
		log.WithError(err).Panic("unable to initialise log")
	}
}

func initializeDB(ctx context.Context) {
	connector := &db.Connection{}
	connector.ConnectMongo(ctx)

	dbase := connector.DB.Database("OMS")
	userCollection := dbase.Collection("orders")
	service.SetupOrderCollection(userCollection)
}

func initializeKafkaProducer(ctx context.Context) {
	kafkaBrokers := config.GetStringSlice(ctx, "onlineKafka.brokers")
	kafkaClientID := config.GetString(ctx, "onlineKafka.clientId")
	kafkaVersion := config.GetString(ctx, "onlineKafka.version")

	producer := kafka.NewProducer(
		kafka.WithBrokers(kafkaBrokers),
		kafka.WithClientID(kafkaClientID),
		kafka.WithKafkaVersion(kafkaVersion),
	)

	log.Printf("Initialized Kafka Producer")
	kafka_producer.Set(producer)
}
