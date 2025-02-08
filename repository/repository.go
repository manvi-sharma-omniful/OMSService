package repository

import (
	"context"
	"time"

	"awesomeProject/Project/OMS/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderCollection struct {
	Collection *mongo.Collection
}

func (C *OrderCollection) Create(order domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := C.Collection.InsertOne(ctx, order)
	if err != nil {
		return err
	}
	return nil
}
