package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserRegister(router *gin.RouterGroup) {
	router.POST("/signUp", SignUp)
	router.GET("/login", Login)
	router.GET("/logout", Logout)
}

func SignUp(c *gin.Context) {
	size_body, error := strconv.Atoi(c.Request.Header.Get("Content-Length"))
	if error != nil {
		c.JSON(http.StatusBadRequest, "error in header Content-Length")
		return
	}
	buf := make([]byte, size_body)
	num, _ := c.Request.Body.Read(buf)
	var user = &User{}
	error = json.Unmarshal(buf[:num], user)
	if error != nil {
		c.JSON(http.StatusBadRequest, "can't desirealize json")
		return
	}
	user.SetPassword(user.PasswordHash)
	error = SaveUser(user)
	if error != nil {
		c.JSON(http.StatusBadRequest, "can't save that object")
		return
	}
	c.JSON(http.StatusCreated, "success")
}

func Login(c *gin.Context) {

}

func Logout(c *gin.Context) {

}
