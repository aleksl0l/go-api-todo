package common

import (
	"fmt"
	"github.com/globalsign/mgo"
)

type Database struct {
	*mgo.Database
}

var DB *mgo.Database

func Init() *mgo.Database {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("db error: ", err)
	}
	DB = session.DB("todo")
	return DB
}

func GetDb() *mgo.Database {
	return DB
}
