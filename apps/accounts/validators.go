package accounts

import "errors"

type UserValidator struct {
	Username string `form:"username" json:"username" binding:"required,alphanum,min=4,max=255"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password"`
	Bio      string `json:"bio" binding:"max=1024"`
	Image    string `json:"image" binding:"omitempty,url"`
}

func (uv *UserValidator) GetData(u *User) error {
	u.Username = uv.Email
	u.Email = uv.Email
	u.Bio = uv.Bio
	if uv.Password == "" && u.Password == "" {
		return errors.New("password is required")
	}
	if err := u.SetPassword(uv.Password); err != nil {
		return err
	}
	if uv.Image != "" {
		u.Image = &uv.Image
	}
	return nil
}


type LoginValidator struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}
