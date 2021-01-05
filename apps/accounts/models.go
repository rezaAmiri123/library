package accounts

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/rezaAmiri123/library/conf"
	"golang.org/x/crypto/bcrypt"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Follow{})
}

type User struct {
	gorm.Model
	Username      string  `gorm:"column:username;unique_index"`
	Password      string  `gorm:"column:password;not null"`
	Email         string  `gorm:"column:email,unique_index"`
	Bio           string  `gorm:"column:bio;size:1024"`
	Image         *string `gorm:"column:image"`
	EmailVerified bool    `gorm:"column:email_verified,default:false"`
	IsActive      bool    `gorm:"column:is_active,default:false"`
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

func (u *User) Update(data interface{}) error {
	db := conf.GetDB()
	return db.Model(u).Update(data).Error
}

type Follow struct {
	gorm.Model
	Follower   User
	FollowerID uint
	Followed   User
	FollowedID uint
}

func (u User) Follow(ou User) error {
	db := conf.GetDB()
	var f Follow
	err := db.FirstOrCreate(&f, Follow{
		FollowerID: u.ID,
		FollowedID: ou.ID,
	}).Error
	return err
}

func (u User) IsFollowed(ou User) bool {
	db := conf.GetDB()
	var f Follow
	db.Where(Follow{
		FollowerID: u.ID,
		FollowedID: ou.ID,
	}).First(&f)
	return f.ID != 0
}

func (u User) UnFollow(ou User) error {
	db := conf.GetDB()
	err := db.Where(Follow{
		FollowerID: u.ID,
		FollowedID: ou.ID,
	}).Delete(Follow{}).Error
	return err
}

func (u User) GetFollowings() []User {
	db := conf.GetDB()
	var fs []Follow
	var fus []User
	tx := db.Begin()
	tx.Where(Follow{FollowerID: u.ID}).Find(&fs)
	for _, f := range fs {
		var fu User
		tx.Model(&f).Related(&fu, "Followed")
		fus = append(fus, fu)
	}
	tx.Commit()
	return fus
}
