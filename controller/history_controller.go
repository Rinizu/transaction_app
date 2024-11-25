package controller

import (
	"net/http"
	"strconv"
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

func (h *HistoryController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/history", h.GetAllHistory)
	router.GET("/history/customer", h.GetHistoryByCustomerID)
}

func (h *HistoryController) GetAllHistory(ctx *gin.Context) {
	histories, err := h.historyUC.GetAllHistory()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "All transaction fetched successfully",
		"data":    histories,
	})
}

func (h *HistoryController) GetHistoryByCustomerID(ctx *gin.Context) {
	customerIDStr := ctx.DefaultQuery("customer_id", "")
	if customerIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "customer_id is required"})
		return
	}

	customerID, err := strconv.Atoi(customerIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid customer_id",
		})
		return
	}

	histories, err := h.historyUC.GetHistoryByCustomerID(customerID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Transaction fetched successfully",
		"data":    histories,
	})
}
