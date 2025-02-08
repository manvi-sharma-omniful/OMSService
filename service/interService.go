package service

import (
	"time"

	interservice_client "github.com/omniful/go_commons/interservice-client"
)

func CreateInterserviceClient() (*interservice_client.Client, error) {
	conn := interservice_client.HTTPTransport()
	config := &interservice_client.Config{
		ServiceName: "InventoryValidation",
		BaseURL:     "http://localhost:8081/inventory/",
		Timeout:     30 * time.Second,
		Transport:   conn,
	}
	return interservice_client.NewClientWithConfig(*config)
}
