package services

import (
	rest_err "github.com/DavibernardesA/CRUD-Go/src/configuration/api_err"
	"github.com/DavibernardesA/CRUD-Go/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(
	userId string) *rest_err.ApiErr {

	logger.Info("Init deleteUser model.",
		zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "deleteUser"))
		return err
	}

	logger.Info(
		"deleteUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))
	return nil
}
