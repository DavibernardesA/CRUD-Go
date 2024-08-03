package services

import (
	rest_err "github.com/DavibernardesA/CRUD-Go/src/configuration/api_err"
	"github.com/DavibernardesA/CRUD-Go/src/model"
	"github.com/DavibernardesA/CRUD-Go/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (
		model.UserDomainInterface, *rest_err.ApiErr)

	FindUserByIDServices(
		id string,
	) (model.UserDomainInterface, *rest_err.ApiErr)
	FindUserByEmailServices(
		email string,
	) (model.UserDomainInterface, *rest_err.ApiErr)

	UpdateUser(string, model.UserDomainInterface) *rest_err.ApiErr
	DeleteUser(string) *rest_err.ApiErr

	LoginUserServices(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, string, *rest_err.ApiErr)
}
