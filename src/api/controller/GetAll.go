package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (sc *sellerController) GetAll(c *gin.Context) {
	response, err := sc.SellerRepository.GetAll()

	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, response)
}
