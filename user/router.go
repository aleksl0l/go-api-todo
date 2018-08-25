package user

import (
	"github.com/gin-gonic/gin"
	"go-api-todo/common"
	"net/http"
)

func UserRegister(router *gin.RouterGroup) {
	router.POST("/signUp", SignUp)
	router.GET("/login", Login)
	router.GET("/logout", Logout)
}

func SignUp(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	user := userModelValidator.userModel
	if err := SaveUser(user); err != nil {
		c.JSON(http.StatusBadRequest, "can't save that object")
		return
	}
	c.JSON(http.StatusCreated, "success")
}

func Login(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	user := userModelValidator.userModel

	if _, err := GetUserByUsername(user.Username); err != nil {
		c.JSON(http.StatusBadRequest, "user not found")
		return
	}
	token := common.GenToken(user.Username)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Logout(c *gin.Context) {

}
