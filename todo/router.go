package todo

import "github.com/gin-gonic/gin"

func TodoRegister(router *gin.RouterGroup) {
	router.GET("/", TodoRetrive)
	router.POST("/", TodoAdd)
	router.DELETE("/", TodoDelete)
}

func TodoRetrive(c *gin.Context) {

}

func TodoAdd(c *gin.Context) {

}

func TodoDelete(c *gin.Context) {

}
