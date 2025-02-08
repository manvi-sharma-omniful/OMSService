package repository

import (
	"awesomeProject/Project/OMS/domain"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type OrderCollection struct {
	Collection *mongo.Collection
}

func (oc *OrderCollection) Create(order domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := oc.Collection.InsertOne(ctx, order)
	if err != nil {
		return err
	}
	return nil
}
