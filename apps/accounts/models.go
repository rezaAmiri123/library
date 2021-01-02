package accounts

import (
	"github.com/jinzhu/gorm"
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
