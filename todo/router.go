package todo

import (
	"github.com/gin-gonic/gin"
	"go-api-todo/common"
	"go-api-todo/user"
	"net/http"
)

func TodoRegister(router *gin.RouterGroup) {
	router.GET("/", TodoRetrive)
	router.POST("/", TodoAdd)
	router.DELETE("/", TodoDelete)
}

func TodoRetrive(c *gin.Context) {
	user_res, _ := c.MustGet("user").(user.User)
	todos, err := GetAllTodoByUser(user_res)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todos})
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "Can't save that object"})
		return
	}
	c.JSON(http.StatusOK, nil)
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
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, nil)
}
