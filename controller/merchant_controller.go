package controller

import (
	"net/http"
	"transaction_app/usecase"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	merchantUC usecase.MerchantUsecase
}

func NewMerchantController(merhantUC usecase.MerchantUsecase) *MerchantController {
	return &MerchantController{
		merchantUC: merhantUC,
	}
}

func (m *MerchantController) RegisterRoutes(router *gin.RouterGroup) {
	api := router.Group("/api")
	{
		api.GET("/merchants", m.GetMerchants)
	}
}

func (m *MerchantController) GetMerchants(ctx *gin.Context) {
	merchants, err := m.merchantUC.GetMerchants()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"merchants": merchants})
}
