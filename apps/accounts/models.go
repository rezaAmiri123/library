package accounts

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
type User struct {
	gorm.Model
	Username      string  `gorm:"column:username;unique_index"`
	Password      string  `gorm:"column:password;not null"`
	Email         string  `gorm:"column:email,unique_index"`
	Bio           string  `gorm:"column:bio;size:1024"`
	Image         *string `gorm:"column:image"`
	EmailVerified bool    `gorm:"column:email_verified"`
	IsActive      bool    `gorm:"column:is_active"`
}

func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}

func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}