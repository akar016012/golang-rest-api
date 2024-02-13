package contoller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context) {
	c.IndentedJSON(http.StatusOK,gin.H{"message":"connected 🛜"})
}