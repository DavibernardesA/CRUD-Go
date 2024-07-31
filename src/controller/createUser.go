package controller

import (
	"fmt"
	"log"

	"github.com/DavibernardesA/CRUD-Go/src/configuration/validation"
	"github.com/DavibernardesA/CRUD-Go/src/controller/model/request"
	"github.com/DavibernardesA/CRUD-Go/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, erro=%s\n", err.Error())
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		Id:    "1",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(200, response)
}
