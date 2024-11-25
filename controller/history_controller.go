package controller

import (
	"net/http"
	"transaction_app/usecase"

	"github.com/gin-gonic/gin"
)

type HistoryController struct {
	historyUC usecase.HistoryUsecase
}

func NewHistoryController(historyUC usecase.HistoryUsecase) *HistoryController {
	return &HistoryController{
		historyUC: historyUC,
	}
}

func (h *HistoryController) RegisterRoutes(router *gin.Engine) {
	router.POST("/transaction", h.CreateTransaction)
}

func (h *HistoryController) CreateTransaction(ctx *gin.Context) {
	var req struct {
		CustomerID int     `json:"customer_id" binding:"required"`
		MerchantID int     `json:"merchant_id" binding:"required"`
		Amount     float64 `json:"amount" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.historyUC.CreateTransaction(req.CustomerID, req.MerchantID, req.Amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "transaction created successfully"})
}
