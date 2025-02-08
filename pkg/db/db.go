package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoConnection interface {
	ConnectMongo(context.Context)
}

type Connection struct {
	DB *mongo.Client
}

func getDatabaseUri() string {
	return "mongodb://127.0.0.1:27017"
}

func (C *Connection) ConnectMongo(con context.Context) {
	fmt.Println("Connecting to mongo...")
	ctx, cancel := context.WithTimeout(con, 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(getDatabaseUri())
	var err error
	C.DB, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

}
