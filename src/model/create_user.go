package model

import (
	"fmt"

	rest_err "github.com/DavibernardesA/CRUD-Go/src/configuration/api_err"
	"github.com/DavibernardesA/CRUD-Go/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.ApiErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
