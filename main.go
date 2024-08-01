package main

import (
	"log"

	"github.com/DavibernardesA/CRUD-Go/src/configuration/logger"
	"github.com/DavibernardesA/CRUD-Go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application.")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
