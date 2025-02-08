package controllers

import (
	"awesomeProject/Project/OMS/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prateek-srivastav-omniful/oms-service/models"
	"github.com/prateek-srivastav-omniful/oms-service/services"
)

func validateCSVFilePath(filePath string) error {
	// Extension check
	if !strings.HasSuffix(filePath, ".csv") {
		return fmt.Errorf("invalid file: not a CSV file")
	}

	// Existance check
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	if err != nil {
		return fmt.Errorf("error accessing file: %v", err)
	}

	// regular file
	if info.IsDir() {
		return fmt.Errorf("invalid file: path points to a directory, not a CSV file")
	}

	return nil
}

// CreateOrder handles incoming order requests, validates them, and stores them in the database

func CreateOrder(c *gin.Context) {
	var orderReq models.OrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		log.Fatal("Error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return
	}
	//
	if err := validateCSVFilePath(orderReq.Path); err != nil {
		log.Fatal("Invalid Path")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := service.ConvertControllerRequestToService(c, orderReq.Path)
	if err != nil {
		log.Fatal("Error", err)
	}

}
