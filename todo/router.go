package todo

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	size_body, error := strconv.Atoi(c.Request.Header.Get("Content-Length"))
	if error != nil {
		c.JSON(http.StatusBadRequest, "error in header Content-Length")
		return
	}
	buf := make([]byte, size_body)
	num, _ := c.Request.Body.Read(buf)
	var todo = &Todo{}
	error = json.Unmarshal(buf[0:num], todo)

	error = DeleteTodo(todo.Name)
	//if error != nil {
	//	c.JSON(http.StatusNotFound, "can't delete todo with name")
	//	return
	//}
	panic("blya")
	c.JSON(http.StatusOK, "success")
}
