package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pedro-phd/iservice-backend/src/api"
	"github.com/pedro-phd/iservice-backend/src/api/controller"
	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	"github.com/pedro-phd/iservice-backend/src/infra/db"
	"github.com/pedro-phd/iservice-backend/src/infra/repositories/sellerRepository"
)

func main() {

	//init env's
	logger.Info("Init application", "main")
	if err := godotenv.Load("../.env"); err != nil {
		logger.Error("Error to load env file", "main", err)
		log.Fatal("Error loading .env file")
	}

	//init dependency
	logger.Info("Init dependency", "main")

	db := db.NewDBConnection()
	repo := sellerRepository.NewSellerRepository(db)

	// seller := &dto.SellerRequest{
	// Name:  "pedro-editado",
	// Email: "pedro-editado@pedro.com",
	// Document:   "123",
	// Phone:      "123",
	// Password:   "123",
	// Image:      "link.com",
	// State:      "SP",
	// City: "Atibaia",
	// Street:     "Rua 1",
	// Number:     "123a",
	// Categories: []int64{1, 2, 3},
	// }

	// repo.Update(5, seller)
	// repo.Create(seller)
	// repo.Delete(9)
	// repo.FindAll()
	// repo.FindById(5)

	//init controller

	router := gin.Default()
	rg := router.Group("/seller")

	sellerController := controller.NewSellerController(repo)
	api.InitRoutesSeller(rg, sellerController)

	if err := router.Run(":3333"); err != nil {
		logger.Error("Api don't have conditions for turns on", "Init api", err)
		panic("Api don't have conditions for turns on")
	}

	//test

}
