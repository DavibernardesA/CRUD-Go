package routes

import (
	"github.com/DavibernardesA/CRUD-Go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:userId", controller.GetUserById)
	r.GET("/user/:userEmail", controller.GetUserByEmail)
	r.POST("/user/:userId", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)
}