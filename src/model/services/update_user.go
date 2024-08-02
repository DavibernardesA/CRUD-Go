package services

import (
	rest_err "github.com/DavibernardesA/CRUD-Go/src/configuration/api_err"
	"github.com/DavibernardesA/CRUD-Go/src/model"
)

func (ud *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.ApiErr {
	return nil
}
