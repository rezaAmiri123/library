package accounts

import "github.com/gin-gonic/gin"

func GetUser(ctx *gin.Context) User {
	u := ctx.MustGet("user").(User)
	return u
}
