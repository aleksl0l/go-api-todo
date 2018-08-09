package todo

import "todoapi/common"

type Todo struct {
	Name string `json:"name" bson:"name"`
	Desc string `json:"description", bson:"description"`
}

func SaveTodo(data interface{}) error {
	db := common.GetDb().DB("todo")
	err := db.C("todo").Insert(data)
	return err
}
