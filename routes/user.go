package routes

import (
	"github.com/akar016012/golang-rest-api/contoller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", contoller.GetUsers)
	router.POST("/", contoller.CreateUser)
	router.DELETE("/:id", contoller.DeleteUser)
	router.PUT("/:id", contoller.UpdateUser)
}