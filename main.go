package main

import (
	"todoapi/common"
	"todoapi/todo"
)

func main() {
	session := common.Init()
	defer session.Close()
	todo.SaveTodo(todo.Todo{
		Name: "task2",
		Desc: "Send otchet",
	})
}
