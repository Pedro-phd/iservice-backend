package controller

import (
	"github.com/gin-gonic/gin"
)

// interface
type ISellerController interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
}

// class
type sellerController struct {
}

func NewSellerController() ISellerController {
	return &sellerController{}
}
