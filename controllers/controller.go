package controllers

import (
	"awesomeProject/Project/OMS/domain"
	"awesomeProject/Project/OMS/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

func validateCSVFilePath(filePath string) error {
	// Extension check
	if !strings.HasSuffix(filePath, ".csv") {
		return fmt.Errorf("invalid file: not a CSV file")
	}

	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	if err != nil {
		return fmt.Errorf("error accessing file: %v", err)
	}

	if info.IsDir() {
		return fmt.Errorf("invalid file: path points to a directory, not a CSV file")
	}
	return nil
}

func CreateOrder(c *gin.Context) {
	var orderReq domain.CreateOrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		log.Fatal("Error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := validateCSVFilePath(orderReq.Path); err != nil {
		log.Fatal("Invalid Path")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// Publishing the path to sqs
	// message:=createMessage(orderReq.Path)
	err := service.ConvertControllerRequestToService(c, orderReq.Path)
	if err != nil {
		log.Fatal("Error", err)
	}
}
