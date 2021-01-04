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
	uv := NewUserValidator()
	if err := ctx.ShouldBind(&uv); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewError("detail", err))
		return
	}
	u, err := uv.Convert()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewValidatorError(err))
		return
	}
	if err := common.SaveObject(u); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewError("detail", err))
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"user": uv})
}

func UserLogin(ctx *gin.Context) {
	lv := NewLoginValidator()
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
