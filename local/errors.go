package local

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Error struct {
	StatusCode int    `json:"-"`
	Origin     string `json:"origin"`
	Code       string `json:"code"`
	Data       any    `json:"data,omitempty"`

	From error `json:"-"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d - %s.%s\nAdditionnal data:\n%vWrapped error:\n%v", e.StatusCode, e.Origin, e.Code, e.Data, e.From)
}

func (e *Error) Apply(c *gin.Context) {
	if e == nil {
		return
	}

	_ = c.Error(e)
	c.JSON(e.StatusCode, e)
	c.Abort()
}
