package service

import (
	"awesomeProject/Project/OMS/domain"
	"context"
	"fmt"
	"github.com/omniful/go_commons/csv"
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
			quantity := record[3]

			quantity, err := strconv.Atoi(quantity)
			if err != nil {
				return nil, fmt.Errorf("invalid quantity %s: %v", quantity, err)
			}

			orderKey := fmt.Sprintf("%s-%s", orderNo, customerName)
			order, exists := orderGroups[orderKey]
			if !exists {
				order = &domain.Order{
					ID:         orderNo,
					UserName:   customerName,
					OrderItems: []domain.OrderItem{},
					Status:     "on_hold",
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}
				orderGroups[orderKey] = order
			}

			skuID, _ = strconv.Atoi(record[4])

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
