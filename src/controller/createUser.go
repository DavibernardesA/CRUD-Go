package controller

import (
	"net/http"

	"github.com/DavibernardesA/CRUD-Go/src/configuration/logger"
	"github.com/DavibernardesA/CRUD-Go/src/configuration/validation"
	"github.com/DavibernardesA/CRUD-Go/src/controller/model/request"
	"github.com/DavibernardesA/CRUD-Go/src/model"
	"github.com/DavibernardesA/CRUD-Go/src/model/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller.", zap.String("journey", "CreateUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info.", err, zap.String("journey", "CreateUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	service := services.NewUserDomainService()
	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully.", zap.String("journey", "CreateUser"))

	c.String(http.StatusOK, "")
}
