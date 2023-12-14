package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	"github.com/pedro-phd/iservice-backend/src/configuration/validation"
	"github.com/pedro-phd/iservice-backend/src/domain/dto"
	"github.com/pedro-phd/iservice-backend/src/domain/models"
)

func (sc *sellerController) Create(c *gin.Context) {

	var sellerRequest dto.SellerRequestDTO

	if err := c.ShouldBindJSON(&sellerRequest); err != nil {
		logger.Error("Error trying to validate user info", "CreateSeller", err)

		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	seller := models.NewSeller(sellerRequest.Name, sellerRequest.Email, sellerRequest.Password, sellerRequest.Products)
	seller.EncryptPassword()

	id, err := sc.SellerRepository.Create(seller)

	if err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	c.JSON(http.StatusCreated, id)
}
