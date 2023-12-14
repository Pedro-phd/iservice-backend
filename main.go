package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pedro-phd/iservice-backend/src/api"
	"github.com/pedro-phd/iservice-backend/src/api/controller"
	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	sellerrepository "github.com/pedro-phd/iservice-backend/src/infra/repository/sellerRepository"
)

func main() {

	//init env's
	logger.Info("Init application", "main")
	if err := godotenv.Load(); err != nil {
		logger.Error("Error to load env file", "main", err)
		log.Fatal("Error loading .env file")
	}

	//init dependency
	logger.Info("Init dependency", "main")
	sellerRepository := sellerrepository.NewSellerRepository()
	sellerController := controller.NewSellercontroller(sellerRepository)

	//init controller

	router := gin.Default()
	// router.Group("/seller")
	api.InitRoutesSeller(&router.RouterGroup, sellerController)

	if err := router.Run(":3333"); err != nil {
		logger.Error("Api don't have conditions for turns on", "Init api", err)
		panic("Api don't have conditions for turns on")
	}

	//test

	// test.GetAll()
}
