package repository

import (
	rest_err "github.com/DavibernardesA/CRUD-Go/src/configuration/api_err"
	"github.com/DavibernardesA/CRUD-Go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER = "MONGODB_USER"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.ApiErr)

	UpdateUser(
		userId string,
		userDomain model.UserDomainInterface,
	) *rest_err.ApiErr

	DeleteUser(
		userId string,
	) *rest_err.ApiErr

	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.ApiErr)
	FindUserByEmailAndPassword(
		email string,
		password string,
	) (model.UserDomainInterface, *rest_err.ApiErr)
	FindUserByID(
		id string,
	) (model.UserDomainInterface, *rest_err.ApiErr)
}
