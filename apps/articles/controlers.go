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
	rg.DELETE("/article/:slug", ArticleDelete)
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

func ArticleDelete(ctx *gin.Context) {
	u := accounts.GetUser(ctx)
	slug := ctx.Param("slug")
	if err := common.DeleteObject(Article{}, &Article{Slug: slug, AuthorID: u.ID}); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
