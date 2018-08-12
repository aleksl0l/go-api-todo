package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TodoRegister(router *gin.RouterGroup) {
	router.GET("/", TodoRetrive)
	router.POST("/", TodoAdd)
	router.DELETE("/", TodoDelete)
}

func TodoRetrive(c *gin.Context) {
	todos, err := GetAllTodo()
	if err != nil {
		c.JSON(http.StatusNotFound, "pakedava")
	}
	c.JSON(http.StatusOK, gin.H{"todos": todos})

}

func TodoAdd(c *gin.Context) {

}

func TodoDelete(c *gin.Context) {

}
