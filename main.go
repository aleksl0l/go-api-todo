package main

import (
	"github.com/gin-gonic/gin"
	"todoapi/common"
	"todoapi/todo"
)

func main() {
	db := common.Init()
	defer db.Session.Close()
	//todo.SaveTodo(todo.Todo{
	//	Name: "task5",
	//	Desc: "Send otchet",
	//})
	//
	//todos, err := todo.GetAllTodo()
	//if err != nil {
	//	fmt.Println("There is error")
	//} else {
	//	fmt.Println(todos)
	//}

	r := gin.Default()
	v1 := r.Group("/api")
	todo.TodoRegister(v1.Group("/todo"))

	r.Run()
}
