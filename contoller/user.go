package contoller

import (
	"net/http"

	"github.com/akar016012/golang-rest-api/config"
	"github.com/akar016012/golang-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.IndentedJSON(http.StatusOK,&users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.IndentedJSON(http.StatusCreated, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.IndentedJSON(http.StatusOK,&user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id=?",c.Param("id")).Delete(&user)
	c.IndentedJSON(http.StatusOK, &user)
}