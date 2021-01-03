package accounts

import (
	"github.com/gin-gonic/gin"
	"github.com/rezaAmiri123/library/apps/common"
	"github.com/rezaAmiri123/library/conf"
	"net/http"
)

func Register(rg *gin.RouterGroup){
	db := conf.GetDB()
	AutoMigrate(db)
	UserAnonRouter(rg.Group("/users"))
}
func UserAnonRouter(rg *gin.RouterGroup) {
	rg.POST("/register", UserCreate)
	rg.POST("/login", UserLogin)
}

func UserCreate(ctx *gin.Context) {
	uv := NewUserValidator()
	if err := ctx.ShouldBind(&uv); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewError("detail", err))
		return
	}
	u, err := uv.Convert()
	if err != nil{
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
	if err:= ctx.ShouldBind(&lv); err!=nil {
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
