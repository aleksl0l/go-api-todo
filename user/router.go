package user

import "github.com/gin-gonic/gin"

func UserRegister(router *gin.RouterGroup) {
	router.POST("/signUp", SignUp)
	router.GET("/login", Login)
	router.GET("/logout", Logout)
}

func SignUp(c *gin.Context) {

}

func Login(c *gin.Context) {

}

func Logout(c *gin.Context) {

}
