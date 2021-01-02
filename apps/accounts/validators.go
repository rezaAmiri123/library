package accounts

import (
	"github.com/gin-gonic/gin"
	"github.com/rezaAmiri123/library/apps/common"
	"github.com/rezaAmiri123/library/conf"
)

type UserValidator struct {
	UserFields struct {
		Username string `json:"username" binding:"required,alphanum,min=4,max=255"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=255"`
		Bio      string `json:"bio" binding:"max=1024"`
		Image    string `json:"image" binding:"omitempty,url"`
	} `json:"-"`
	User `json:"-"`
}

func (uv *UserValidator)Bind(ctx *gin.Context) error {
	err := common.Bind(ctx, uv)
	if err != nil {
		return err
	}
	uv.User.Username = uv.UserFields.Username
	uv.User.Email = uv.UserFields.Email
	uv.User.Bio = uv.UserFields.Bio
	if uv.UserFields.Password != conf.NBRandomPassword {
		uv.User.SetPassword(uv.UserFields.Password)
	}
	if uv.UserFields.Image != "" {
		uv.User.Image = &uv.UserFields.Image
	}
	return nil
}

func NewUserValidator() UserValidator {
	uv := UserValidator{}
	return uv
}
