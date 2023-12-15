package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	rest_err "github.com/pedro-phd/iservice-backend/src/domain/error"
)

func (sc *sellerController) GetByEmail(c *gin.Context) {

	email, ok := c.Params.Get("email")
	if !ok {
		err := rest_err.NewBadRequestError("email is necessary")
		c.JSON(err.Code, err)
	}

	res, err := sc.SellerRepository.GetByEmail(email)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, res)

}
