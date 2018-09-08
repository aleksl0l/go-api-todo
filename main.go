package main

import (
	"github.com/gin-gonic/gin"
	"go-api-todo/common"
	"go-api-todo/todo"
	"go-api-todo/user"
)

func main() {
	db := common.Init()
	defer db.Session.Close()

	r := gin.Default()

	v1 := r.Group("/api")

	todo.TodoRegister(v1.Group("/todo"))
	user.UserRegister(v1.Group("/user"))

	r.Run()
}
