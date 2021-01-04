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

}
func UserRouter(rg *gin.RouterGroup) {
	rg.POST("/register", UserCreate)
	rg.POST("/login", UserLogin)
	rg.GET("/user", AuthMiddleware(true), UserRetrieve)
}

func UserCreate(ctx *gin.Context) {
	var uv UserValidator
	if err := ctx.ShouldBind(&uv); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewError("detail", err))
		return
	}
	var u User
	if err := uv.GetData(&u); err != nil{
		ctx.JSON(http.StatusBadRequest, common.NewError("detail", err))
		return
	}
	if err := common.SaveObject(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewError("detail", err))
		return
	}
	var us  UserSerializer
	ctx.JSON(http.StatusCreated, gin.H{"user": us.Response(u)})
}

func UserLogin(ctx *gin.Context) {
	var lv LoginValidator
	if err := ctx.ShouldBind(&lv); err != nil {
		ctx.JSON(http.StatusNotFound, common.NewError("detail", err))
		return
	}
	var u User
	if err := common.FindObject(&u, User{Email: lv.Email}); err != nil {
		ctx.JSON(http.StatusNotFound, common.NewError("detail", err))
		return
	}
	if err := u.CheckPassword(lv.Password); err != nil {
		ctx.JSON(http.StatusNotFound, common.NewError("detail", err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": common.GetToken(u.ID)})
}

func UserRetrieve(ctx *gin.Context) {
	u := ctx.MustGet("user").(User)
	var us UserSerializer
	ctx.JSON(http.StatusOK, gin.H{"user": us.Response(u)})
}
