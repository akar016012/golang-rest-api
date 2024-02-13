package routes

import (
	"github.com/akar016012/golang-rest-api/contoller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", contoller.UserController)
}