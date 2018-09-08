package todo

import (
	"github.com/globalsign/mgo/bson"
	"go-api-todo/common"
	"go-api-todo/user"
)

type Todo struct {
	Name   string        `json:"name" bson:"name"`
	Desc   string        `json:"description" bson:"description"`
	UserID bson.ObjectId `json:"-" bson:"_userID"`
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
	return res, err
}

func GetAllTodoByUser(user user.User) ([]Todo, error) {
	db := common.GetDb()
	var res []Todo
	err := db.C("todo").Find(bson.M{"_userID": user.Id}).All(&res)
	return res, err
}

func DeleteTodo(name string) error {
	db := common.GetDb()
	err := db.C("todo").Remove(bson.M{"name": name})
	return err
}
