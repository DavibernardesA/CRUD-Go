package services

import (
	rest_err "github.com/DavibernardesA/CRUD-Go/src/configuration/api_err"
	"github.com/DavibernardesA/CRUD-Go/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.ApiErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.ApiErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.ApiErr)
	DeleteUser(string) *rest_err.ApiErr
}
