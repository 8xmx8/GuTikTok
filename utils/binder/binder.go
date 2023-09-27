package binder

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// Bind Query/Json 双绑定,为后续网页端提供支持
func Bind(c *gin.Context, data any) error {
	if err := c.ShouldBindQuery(data); err != nil {
		if err2 := c.ShouldBindJSON(data); err2 != nil {
			return errors.Join(err, err2)
		}
	}
	return nil
}
