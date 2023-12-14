package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pedro-phd/iservice-backend/src/api/controller"
)

func InitRoutesSeller(r *gin.RouterGroup, controller controller.SellerController) {
	r.GET("/seller/all", controller.GetAll)
	r.POST("/seller/", controller.Create)
}
