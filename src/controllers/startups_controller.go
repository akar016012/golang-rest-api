package controllers

import (
	"net/http"

	"github.com/akar016012/golang-rest-api/src/models"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context){
	c.IndentedJSON(http.StatusOK, gin.H{"TEST": "YOU SCORED 💯"})
}


func GetAllStartups(c *gin.Context) {
 startups, err := models.FetchAllStartups()
 if err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
  return
 }
 c.JSON(http.StatusOK, gin.H{"message": "Startups fetched successfully", "status": "success", "data": startups})
}