package todo

import (
	"fmt"
	"todoapi/common"
)

type Todo struct {
	Name string `json:"name" bson:"name"`
	Desc string `json:"description" bson:"description"`
}

func SaveTodo(data interface{}) error {
	db := common.GetDb()
	err := db.C("todo").Insert(data)
	return err
}

func GetAllTodo() ([]Todo, error) {
	db := common.GetDb()
	var res []Todo
	err := db.C("todo").Find(nil).All(&res)
	fmt.Println(res)
	return res, err
}
