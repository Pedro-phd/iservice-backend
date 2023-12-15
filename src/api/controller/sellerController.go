package controller

import (
	"github.com/gin-gonic/gin"
	sellerrepository "github.com/pedro-phd/iservice-backend/src/infra/repository/sellerRepository"
)

type SellerController interface {
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	GetByEmail(c *gin.Context)
}

type sellerController struct {
	SellerRepository sellerrepository.ISellerRepository
}

func NewSellercontroller(respoistory sellerrepository.ISellerRepository) SellerController {
	return &sellerController{respoistory}
}
