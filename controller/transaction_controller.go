package controller

import (
	"net/http"
	"transaction_app/middleware"
	"transaction_app/usecase"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionUC usecase.HistoryUsecase
}

func NewTransactionController(transactionUC usecase.HistoryUsecase) *TransactionController {
	return &TransactionController{
		transactionUC: transactionUC,
	}
}

func (t *TransactionController) RegisterRoutes(router *gin.RouterGroup) {
	transaction := router.Group("/transaction")
	transaction.Use(middleware.CustomerAuthMiddleware())
	{
		transaction.POST("", t.CreateTransaction)
	}
}

func (t *TransactionController) CreateTransaction(ctx *gin.Context) {
	customerID, exists := ctx.Get("customer_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Customer ID not found",
		})
		return
	}

	var req struct {
		CustomerID int     `json:"customer_id" binding:"required"`
		MerchantID int     `json:"merchant_id" binding:"required"`
		Amount     float64 `json:"amount" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if req.CustomerID != customerID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized to perform this transaction",
		})
		return
	}

	if req.Amount <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "amount must be greater than 0",
		})
		return
	}

	err := t.transactionUC.CreateTransaction(req.CustomerID, req.MerchantID, float64(req.Amount))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Transaction successfully",
	})
}
