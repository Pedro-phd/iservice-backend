package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (sc *sellerController) GetById(c *gin.Context) {

	// id, ok := c.Params.Get("id")
	// if !ok {
	// 	err := rest_err.NewBadRequestError("id is necessary")
	// 	c.JSON(err.Code, err)
	// }

	// res, err := sc.SellerRepository.GetById(id)

	// if err != nil {
	// 	c.JSON(err.Code, err)
	// 	return
	// }

	c.JSON(http.StatusOK, nil)

}
