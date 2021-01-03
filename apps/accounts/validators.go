package accounts

type UserValidator struct {
	Username string `form:"username" json:"username" binding:"required,alphanum,min=4,max=255"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
	Bio      string `json:"bio" binding:"max=1024"`
	Image    string `json:"image" binding:"omitempty,url"`
}

func (uv *UserValidator) Convert() (*User, error) {
	var u User
	u.Username = uv.Email
	u.Email = uv.Email
	u.Bio = uv.Bio
	if err := u.SetPassword(uv.Password); err != nil {
		return &User{}, err
	}
	if uv.Image != "" {
		u.Image = &uv.Image
	}
	return &u, nil
}

func NewUserValidator() UserValidator {
	uv := UserValidator{}
	return uv
}

type LoginValidator struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}

func NewLoginValidator () LoginValidator {
	return LoginValidator{}
}
