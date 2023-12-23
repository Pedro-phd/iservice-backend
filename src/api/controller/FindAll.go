package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	rest_err "github.com/pedro-phd/iservice-backend/src/domain/error"
)

func (sc *sellerController) FindAll(c *gin.Context) {
	sellers, err := sc.repo.FindAll()
	if err != nil {
		res_err := rest_err.NewBadRequestError(err.Error())
		c.JSON(res_err.Code, res_err)
		return
	}

	c.JSON(http.StatusOK, sellers)
}
