package accounts

type UserSerializer struct {
	ID       uint    `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
}

func (us *UserSerializer) Response(u User) UserSerializer{
	us.ID = u.ID
	us.Email = u.Email
	us.Username = u.Username
	us.Image = u.Image
	us.Bio = u.Bio
	return *us
}
