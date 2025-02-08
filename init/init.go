package init

import (
	"awesomeProject/Project/OMS/pkg/db"
	"awesomeProject/Project/OMS/service"
	"context"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/log"
)

func Initialize(ctx context.Context) {
	initializeLog(ctx)
	initializeDB(ctx)
	initialiseSQSProducer(ctx)
	initialiseSQSConsumer(ctx)
}

func initialiseSQSConsumer(ctx context.Context) {

}

func initialiseSQSProducer(ctx context.Context) {

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
