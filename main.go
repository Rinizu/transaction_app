package main

import (
	"log"
	"transaction_app/config"
	"transaction_app/controller"
	"transaction_app/repository"
	"transaction_app/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	customerRepo := repository.NewCustomerRepository(cfg.CustomerFile)
	merchantRepo := repository.NewMerchantRepository(cfg.MerchantFile)
	historyRepo := repository.NewHistoryRepository(cfg.HistoryFile)

	customerUC := usecase.NewCustomerUsecase(customerRepo)
	merchantUC := usecase.NewMerchantUsecase(merchantRepo)
	historyUC := usecase.NewHistoryUsecase(customerRepo, merchantRepo, historyRepo)

	customerController := controller.NewCustomerController(customerUC)
	merchantController := controller.NewMerchantController(merchantUC)
	historyController := controller.NewHistoryController(historyUC)

	router := gin.Default()
	api := router.Group("/api")
	{
		customerController.RegisterRoutes(api)
		merchantController.RegisterRoutes(api)
		historyController.RegisterRoutes(api)
	}

	log.Printf("Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
