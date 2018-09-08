package todo

import (
	"github.com/gin-gonic/gin"
	"go-api-todo/common"
	"gopkg.in/mgo.v2/bson"
)

type TodoModelValidator struct {
	Name      string        `form:"username" json:"name" binding:"exists,alphanum,max=255"`
	Desc      string        `form:"password" json:"description" binding:"exists,max=511"`
	UserID    bson.ObjectId `form:"-" json:"-"`
	todoModel Todo          `json:"-"`
}

func (self *TodoModelValidator) Bind(c *gin.Context) error {

	if err := common.Bind(c, self); err != nil {
		return err
	}
	self.todoModel.Name = self.Name
	self.todoModel.Desc = self.Desc
	return nil
}

func NewTodoModelValidator() TodoModelValidator {
	TodoModelValidator := TodoModelValidator{}
	return TodoModelValidator
}
