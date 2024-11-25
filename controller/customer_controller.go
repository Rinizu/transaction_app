package controller

import (
	"net/http"
	"transaction_app/usecase"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerUC usecase.CustomerUsecase
}

func NewCustomerController(customerUC usecase.CustomerUsecase) *CustomerController {
	return &CustomerController{
		customerUC: customerUC,
	}
}

func (c *CustomerController) RegisterRoutes(router *gin.RouterGroup) {
	customer := router.Group("/customer")
	{
		customer.POST("/register", c.RegisterCustomer)
		customer.POST("/login", c.LoginCustomer)
		customer.POST("/logout", c.LogoutCustomer)
	}
}

func (c *CustomerController) RegisterCustomer(ctx *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "name, email and password are required"})
		return
	}

	if err := c.customerUC.RegisterCustomer(req.Name, req.Email, req.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "customer registered successfully"})
}

func (c *CustomerController) LoginCustomer(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.customerUC.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("token", token, 3600, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "customer login successfully",
		"token":   token,
	})
}

func (c *CustomerController) LogoutCustomer(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "customer logout successfully",
	})
}
