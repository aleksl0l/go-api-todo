package todo

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-api-todo/user"
	"net/http"
	"strconv"
)

func TodoRegister(router *gin.RouterGroup) {
	router.GET("/", user.JWTAuthorization(), TodoRetrive)
	router.POST("/", user.JWTAuthorization(), TodoAdd)
	router.DELETE("/", user.JWTAuthorization(), TodoDelete)
}

func TodoRetrive(c *gin.Context) {
	todos, err := GetAllTodo()
	if err != nil {
		c.JSON(http.StatusNotFound, "pakedava")
	}
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func TodoAdd(c *gin.Context) {
	size_body, error := strconv.Atoi(c.Request.Header.Get("Content-Length"))
	if error != nil {
		c.JSON(http.StatusBadRequest, "error in header Content-Length")
		return
	}
	buf := make([]byte, size_body)
	num, _ := c.Request.Body.Read(buf)
	var todo = &Todo{}
	error = json.Unmarshal(buf[0:num], todo)
	if error != nil {
		c.JSON(http.StatusBadRequest, "can't deserialize json")
		return
	}
	error = SaveTodo(todo)
	if error != nil {
		c.JSON(http.StatusBadRequest, "can't save that object")
		return
	}
	c.JSON(http.StatusOK, "success")
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
	if error != nil {
		c.JSON(http.StatusNotFound, "can't delete todo with name")
		return
	}
	c.JSON(http.StatusOK, "success")
}
