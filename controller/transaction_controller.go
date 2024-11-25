package controller

import (
	"net/http"
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
	router.POST("/transaction", t.CreateTransaction)
}

func (t *TransactionController) CreateTransaction(ctx *gin.Context) {
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
