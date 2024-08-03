package main

import (
	"context"
	"log"

	mongodb "github.com/DavibernardesA/CRUD-Go/src/configuration/database/mongoDB"
	"github.com/DavibernardesA/CRUD-Go/src/configuration/logger"
	"github.com/DavibernardesA/CRUD-Go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")

	godotenv.Load()

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
