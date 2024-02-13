package main

import (
	"github.com/akar016012/golang-rest-api/config"
	"github.com/akar016012/golang-rest-api/routes"
	"github.com/gin-gonic/gin"
)


func main(){
router := gin.New()
config.Connect()
routes.UserRoute(router)
router.Run(":8080")
}