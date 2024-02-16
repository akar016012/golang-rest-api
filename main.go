package main

import (
	"github.com/akar016012/golang-rest-api/src/models"
	"github.com/akar016012/golang-rest-api/src/routes"
	"github.com/akar016012/golang-rest-api/src/utils"
)


func main() {
utils.LoadEnv()
models.OpenDatabaseConnection()
models.AutoMigrateModels()
router := routes.SetupRoutes()
router.Run(":8080")	
}