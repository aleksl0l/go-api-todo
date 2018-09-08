package todo

import (
	"github.com/gin-gonic/gin"
	"go-api-todo/common"
	"go-api-todo/user"
	"net/http"
)

func TodoRegister(router *gin.RouterGroup) {
	router.GET("/", user.JWTAuthorization(), TodoRetrive)
	router.POST("/", user.JWTAuthorization(), TodoAdd)
	router.DELETE("/", user.JWTAuthorization(), TodoDelete)
}

func TodoRetrive(c *gin.Context) {
	todos, err := GetAllTodo()
	if err != nil {
		c.JSON(http.StatusNotFound, "pakedava")
	}
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func TodoAdd(c *gin.Context) {
	todoModelValidator := NewTodoModelValidator()
	if err := todoModelValidator.Bind(c); err != nil {
		common.RenderResponse(c, http.StatusUnprocessableEntity, common.NewValidatorError(err), nil)
		return
	}
	todo := todoModelValidator.todoModel

	user_res, _ := c.MustGet("user").(user.User)
	todo.UserID = user_res.Id
	err := SaveTodo(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, "can't save that object")
		return
	}
	c.JSON(http.StatusOK, "success")
}

func TodoDelete(c *gin.Context) {
	todoModelValidator := NewTodoModelValidator()
	if err := todoModelValidator.Bind(c); err != nil {
		common.RenderResponse(c, http.StatusUnprocessableEntity, common.NewValidatorError(err), nil)
		return
	}
	todo := todoModelValidator.todoModel

	err := DeleteTodo(todo.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, "can't delete todo with name")
		return
	}
	c.JSON(http.StatusOK, "success")
}
