package service

import (
	"awesomeProject/Project/OMS/domain"
	"context"
	"fmt"
	"github.com/omniful/go_commons/csv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"strconv"
	"time"
)

func CSVOperation(filePath string) ([]*domain.Order, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	orderGroups := make(map[string]*domain.Order)
	CSV, err := csv.NewCommonCSV(
		csv.WithBatchSize(100),
		csv.WithSource(csv.Local),
		csv.WithLocalFileInfo(filePath),
		csv.WithHeaderSanitizers(csv.SanitizeAsterisks, csv.SanitizeToLower),
		csv.WithDataRowSanitizers(csv.SanitizeSpace, csv.SanitizeToLower),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize CSV reader: %v", err)
	}

	err = CSV.InitializeReader(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize CSV reader: %v", err)
	}

	for !CSV.IsEOF() {
		var records csv.Records
		records, err := CSV.ReadNextBatch()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Processing records:")
		fmt.Println(records)
		for _, record := range records {
			fmt.Println(record)
			orderNo := record[0]
			customerName := record[1]
			skuID := record[2]
			quantityStr := record[3]

			quantity, err := strconv.Atoi(quantityStr)
			if err != nil {
				return nil, fmt.Errorf("invalid quantity %s: %v", quantityStr, err)
			}

			orderKey := fmt.Sprintf("%s-%s", orderNo, customerName)
			order, exists := orderGroups[orderKey]
			if !exists {
				now := primitive.NewDateTimeFromTime(time.Now())
				order = &domain.Order{
					ID:         orderNo,
					HubID:      string,
					TenantID:   string,
					UserName:   customerName,
					OrderItems: []domain.OrderItem{},
					Status:     "on_hold",
					CreatedAt:  now,
					UpdatedAt:  now,
				}
				orderGroups[orderKey] = order
			}

			orderItem := domain.OrderItem{
				SKU_ID:   skuID,
				Quantity: quantity,
			}
			order.OrderItems = append(order.OrderItems, orderItem)
		}
	}

	var orders []*domain.Order
	for _, order := range orderGroups {
		orders = append(orders, order)
	}

	return orders, nil
}
