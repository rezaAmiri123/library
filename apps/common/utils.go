package common

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Bind(ctx *gin.Context, obj interface{}) error {
	bind := binding.Default(ctx.Request.Method, ctx.ContentType())
	return ctx.ShouldBindWith(obj, bind)
}
