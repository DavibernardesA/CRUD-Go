package main

import (
	"github.com/DavibernardesA/CRUD-Go/src/controller"
	"github.com/DavibernardesA/CRUD-Go/src/model/repository"
	"github.com/DavibernardesA/CRUD-Go/src/model/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := services.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
