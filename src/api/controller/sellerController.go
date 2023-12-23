package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pedro-phd/iservice-backend/src/infra/repositories/sellerRepository"
)

// interface
type ISellerController interface {
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	Create(c *gin.Context)
	// Delete(c *gin.Context)
	// Update(c *gin.Context)
}

// class
type sellerController struct {
	repo sellerRepository.ISellerRepository
}

func NewSellerController(repo sellerRepository.ISellerRepository) ISellerController {
	return &sellerController{repo}
}
