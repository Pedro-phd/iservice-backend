package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pedro-phd/iservice-backend/src/api/controller"
)

func InitRoutesSeller(r *gin.RouterGroup, controller controller.ISellerController) {
	r.GET("/", controller.FindAll)
	r.GET("/:id", controller.FindById)
	r.POST("/", controller.Create)
	// r.DELETE("/seller/:id", controller.Delete)
}
