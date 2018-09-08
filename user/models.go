package user

import (
	"errors"
	"github.com/globalsign/mgo/bson"
	"go-api-todo/common"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username     string        `json:"username" bson:"username"`
	PasswordHash string        `json:"password" bson:"password"`
}

func SaveUser(data interface{}) error {
	db := common.GetDb()
	err := db.C("user").Insert(data)
	return err
}

//func GetAllUsers() ([]User, error) {
//	db := common.GetDb()
//	var res []User
//	err := db.C("user").Find(nil).All(&res)
//	return res, err
//}

func GetUserByUsername(username string) (User, error) {
	db := common.GetDb()
	var res User
	err := db.C("user").Find(bson.M{"username": username}).One(&res)
	return res, err
}

func (u *User) CheckPassword(password string) bool {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return nil == bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}
