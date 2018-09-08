package common

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"time"
)

const NBSecretPassword = "A String Very Very Very Strong!!@##$!@#$"

//const NBRandomPassword = "A String Very Very Very Niubilty!!@##$!@#4"

func GenToken(id string) string {
	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
	jwt_token.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token, _ := jwt_token.SignedString([]byte(NBSecretPassword))
	return token
}

func KeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(NBSecretPassword), nil
}

func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		if v.Param != "" {
			res.Errors[v.Field] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
		} else {
			res.Errors[v.Field] = fmt.Sprintf("{key: %v}", v.Tag)
		}
	}
	return res
}

func RenderResponse(c *gin.Context, code int, errors interface{}, data interface{}) {
	if errors == nil {
		errors = CommonError{Errors: gin.H{"errors": nil}}
	}
	err := errors.(CommonError)

	body := gin.H{"data": data}
	for k, v := range err.Errors {
		body[k] = v
	}
	c.JSON(code, body)
}
