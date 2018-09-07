package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api-todo/common"
	"go-api-todo/todo"
	"go-api-todo/user"
)

func main() {
	db := common.Init()
	defer db.Session.Close()

	todo.SaveTodo(todo.Todo{
		Name: "task5",
		Desc: "Send otchet",
	})

	//
	//todos, err := todo.GetAllTodo()
	//if err != nil {
	//	fmt.Println("There is error")
	//} else {
	//	fmt.Println(todos)
	//}

	fmt.Println(common.GenToken("test"))

	//u := user.User{}
	//u.SetPassword("puska")
	//u.Username = "puska"
	//fmt.Println(u)
	//fmt.Println(u.CheckPassword("puska1"))
	//
	r := gin.Default()

	v1 := r.Group("/api")

	//authorized := r.Group("/")
	//authorized.Use(Puk())
	//todo.TodoRegister(authorized.Group("/todo"))

	todo.TodoRegister(v1.Group("/todo"))
	user.UserRegister(v1.Group("/user"))

	r.Run()
}
