package common

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type Database struct {
	*mgo.Session
}

var Session *mgo.Session

func Init() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("db error: ", err)
	}
	Session = session
	return Session
}

func GetDb() *mgo.Session {
	return Session
}

//func GetAllTodos() {
//	DB.C("TdoList").Find(interface{});
//}
