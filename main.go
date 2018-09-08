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

	todo_group := v1.Group("/todo")
	todo_group.Use(user.JWTAuthorization())

	todo.TodoRegister(todo_group)
	user.UserRegister(v1.Group("/user"))

	r.Run()
}
