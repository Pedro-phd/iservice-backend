package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	rest_err "github.com/pedro-phd/iservice-backend/src/domain/error"
)

func (sc *sellerController) FindById(c *gin.Context) {

	id, ok := c.Params.Get("id")
	if !ok {
		err := rest_err.NewBadRequestError("id is necessary")
		c.JSON(err.Code, err)
	}
	res, err := sc.repo.FindById(id)

	if err != nil {
		err, ok := err.(*rest_err.RestErr)

		if ok {
			c.JSON(err.Code, err)
			return
		}

		err_res := rest_err.NewBadRequestError(err.Error())
		c.JSON(err_res.Code, err_res)
		return
	}

	c.JSON(http.StatusOK, res)

}
