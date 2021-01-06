package articles

import (
	"github.com/gin-gonic/gin"
	"github.com/rezaAmiri123/library/apps/accounts"
	"github.com/rezaAmiri123/library/apps/common"
	"github.com/rezaAmiri123/library/conf"
	"net/http"
)

func Register(rg *gin.RouterGroup) {
	db := conf.GetDB()
	AutoMigrate(db)
	ArticleRouter(rg.Group("/articles", accounts.AuthMiddleware(true)))
}
func ArticleRouter(rg *gin.RouterGroup) {
	rg.POST("/article", ArticleCreate)
}

func ArticleCreate(ctx *gin.Context) {
	var av ArticleValidator
	if err := ctx.ShouldBind(&av); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}
	var a Article
	if err := av.SetData(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}
	a.Author = accounts.GetUser(ctx)
	if err := common.SaveObject(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}
	var as ArticleSerializer
	ctx.JSON(http.StatusCreated, gin.H{"article": as.Response(a)})
}
