package accounts

import (
	"github.com/gin-gonic/gin"
	"github.com/rezaAmiri123/library/apps/common"
	"github.com/rezaAmiri123/library/conf"
	"net/http"
)

func Register(rg *gin.RouterGroup) {
	db := conf.GetDB()
	AutoMigrate(db)
	UserRouter(rg.Group("/users"))
	ProfileRouter(rg.Group("/profiles", AuthMiddleware(true)))

}
func UserRouter(rg *gin.RouterGroup) {
	rg.POST("/register", UserCreate)
	rg.POST("/login", UserLogin)
	rg.GET("/user", AuthMiddleware(true), UserRetrieve)
	rg.PUT("/user", AuthMiddleware(true), UserUpdate)
}
func ProfileRouter(rg *gin.RouterGroup) {
	rg.GET("/:username", ProfileRetrieve)
	rg.POST("/:username/follow", ProfileFollow)
}

func UserCreate(ctx *gin.Context) {
	var uv UserValidator
	if err := ctx.ShouldBind(&uv); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}
	var u User
	if err := uv.SetData(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}
	if err := common.SaveObject(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}
	var us UserSerializer
	ctx.JSON(http.StatusCreated, gin.H{"user": us.Response(u)})
}

func UserLogin(ctx *gin.Context) {
	var lv LoginValidator
	if err := ctx.ShouldBind(&lv); err != nil {
		ctx.JSON(http.StatusNotFound, common.ErrorResponse(err))
		return
	}
	var u User
	if err := common.FindObject(&u, User{Email: lv.Email}); err != nil {
		ctx.JSON(http.StatusNotFound, common.ErrorResponse(err))
		return
	}
	if err := u.CheckPassword(lv.Password); err != nil {
		ctx.JSON(http.StatusNotFound, common.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": common.GetToken(u.ID)})
}

func UserRetrieve(ctx *gin.Context) {
	u := GetUser(ctx)
	var us UserSerializer
	ctx.JSON(http.StatusOK, gin.H{"user": us.Response(u)})
}

func UserUpdate(ctx *gin.Context) {
	u := GetUser(ctx)
	var uv UserValidator
	if err := ctx.ShouldBind(&uv); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}
	if err := uv.SetData(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}
	if err := common.SaveObject(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}
	var us UserSerializer
	ctx.JSON(http.StatusCreated, gin.H{"user": us.Response(u)})
}

func ProfileRetrieve(ctx *gin.Context) {
	username := ctx.Param("username")
	var u User
	if err := common.FindObject(&u, User{Username: username}); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}
	var ps ProfileSerializer
	ctx.JSON(http.StatusCreated, gin.H{"user": ps.Response(u)})
}

func ProfileFollow(ctx *gin.Context) {
	username := ctx.Param("username")
	var ou User
	if err := common.FindObject(&ou, User{Username: username}); err != nil {
		ctx.JSON(http.StatusNotFound, common.ErrorResponse(err))
		return
	}
	u := GetUser(ctx)
	if err := u.Follow(ou); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err))
		return
	}
	var ps ProfileSerializer
	ctx.JSON(http.StatusCreated, gin.H{"user": ps.Response(ou)})
}
