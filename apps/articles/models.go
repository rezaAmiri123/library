package articles

import (
	"github.com/jinzhu/gorm"
	"github.com/rezaAmiri123/library/apps/accounts"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Favorite{})
}

type Article struct {
	gorm.Model
	Slug        string `gorm:"unique_index"`
	Title       string
	Description string `gorm:"size:2048"`
	Body        string `gorm:"size:2048"`
	Author      accounts.User
	AuthorID    uint
}

type Favorite struct {
	gorm.Model
	User      accounts.User
	UserID    uint
	Article   Article
	ArticleId uint
}
