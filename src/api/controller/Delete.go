package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	rest_err "github.com/pedro-phd/iservice-backend/src/domain/error"
)

func (sc *sellerController) Delete(c *gin.Context) {

	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, "id is necessary")
	}
	result := sc.SellerRepository.Remove(id)

	if !result {
		err := rest_err.NewBadRequestError("Not deleted")
		c.JSON(err.Code, err)
	}

	c.JSON(http.StatusOK, "")
}
