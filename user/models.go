package user

import (
	"errors"
	"go-api-todo/common"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username     string `json:"username",bson:"username"`
	PasswordHash string `json:"password",bson:"password"`
}

func SaveUser(data interface{}) error {
	db := common.GetDb()
	err := db.C("user").Insert(data)
	return err
}

func GetAllUsers() ([]User, error) {
	db := common.GetDb()
	var res []User
	err := db.C("user").Find(nil).All(&res)
	//fmt.Println(res)
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
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}
