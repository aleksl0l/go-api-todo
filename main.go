package main

import (
	"fmt"
	"todoapi/common"
	"todoapi/todo"
)

func main() {
	db := common.Init()
	defer db.Session.Close()
	todo.SaveTodo(todo.Todo{
		Name: "task5",
		Desc: "Send otchet",
	})

	todos, err := todo.GetAllTodo()
	if err != nil {
		fmt.Println("There is error")
	} else {
		fmt.Println(todos)
	}

}
