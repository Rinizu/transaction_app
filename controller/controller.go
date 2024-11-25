package controller

import (
	"transaction_app/usecase"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	CustomerController    *CustomerController
	MerchantController    *MerchantController
	TransactionController *TransactionController
	HistoryController     *HistoryController
}

func NewController(
	customerUC usecase.CustomerUsecase,
	merchantUC usecase.MerchantUsecase,
	historyUC usecase.HistoryUsecase,
) *Controller {
	return &Controller{
		CustomerController:    NewCustomerController(customerUC),
		MerchantController:    NewMerchantController(merchantUC),
		TransactionController: NewTransactionController(historyUC),
		HistoryController:     NewHistoryController(historyUC),
	}
}

func (c *Controller) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		c.CustomerController.RegisterRoutes(api)
		c.MerchantController.RegisterRoutes(api)
		c.TransactionController.RegisterRoutes(api)
		c.HistoryController.RegisterRoutes(api)
	}
}
