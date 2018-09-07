package user

import (
	"github.com/gin-gonic/gin"
	"go-api-todo/common"
	"net/http"
)

func UserRegister(router *gin.RouterGroup) {
	router.POST("/signUp", SignUp)
	router.POST("/login", Login)
	router.GET("/logout", Logout)
}

func SignUp(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		common.RenderResponse(c, http.StatusUnprocessableEntity, common.NewValidatorError(err), nil)
		return
	}
	user := userModelValidator.userModel
	if _, err := GetUserByUsername(user.Username); err == nil {
		common.RenderResponse(c, http.StatusBadRequest, common.CommonError{Errors: gin.H{"errors": "user already exist"}}, nil)
		//c.JSON(http.StatusBadRequest, "user not found")
		return
	}
	if err := SaveUser(user); err != nil {
		common.RenderResponse(c, http.StatusBadRequest, common.CommonError{Errors: gin.H{"errors": "can't save that object"}}, nil)
		return
	}
	c.JSON(http.StatusCreated, "success")
}

func Login(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		common.RenderResponse(c, http.StatusUnprocessableEntity, common.NewValidatorError(err), nil)
		return
	}
	user := userModelValidator.userModel

	if _, err := GetUserByUsername(user.Username); err != nil {
		common.RenderResponse(c, http.StatusBadRequest, common.CommonError{Errors: gin.H{"errors": "user not found"}}, nil)
		//c.JSON(http.StatusBadRequest, "user not found")
		return
	}
	token := common.GenToken(user.Username)
	common.RenderResponse(c, http.StatusOK, common.CommonError{Errors: gin.H{"errors": nil}}, gin.H{"token": token})
}

func Logout(c *gin.Context) {

}
