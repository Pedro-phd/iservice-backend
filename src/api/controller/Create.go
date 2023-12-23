package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	"github.com/pedro-phd/iservice-backend/src/configuration/validation"
	"github.com/pedro-phd/iservice-backend/src/domain/dto"
	rest_err "github.com/pedro-phd/iservice-backend/src/domain/error"
)

type Response struct {
	Id string `json:"id"`
}

func (sc *sellerController) Create(c *gin.Context) {
	journey := "Create - Seller Controller"

	var sellerRequest dto.SellerRequest

	if err := c.ShouldBindJSON(&sellerRequest); err != nil {
		logger.Error("Error trying to validate user info", "CreateSeller", err)

		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	id, err := sc.repo.Create(&sellerRequest)

	if err != nil {
		logger.Error("Error to create seller", journey, err)
		errResponse := rest_err.NewBadRequestError(err.Error())
		c.JSON(errResponse.Code, errResponse)
		return
	}

	reponse := Response{
		Id: id,
	}

	c.JSON(http.StatusCreated, reponse)
}
