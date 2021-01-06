package articles

import (
	"github.com/jinzhu/gorm"
	"github.com/rezaAmiri123/library/apps/accounts"
	"github.com/rezaAmiri123/library/conf"
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

func (a *Article) FavoriteBy(u *accounts.User) error {
	db := conf.GetDB()
	var f Favorite
	err := db.FirstOrCreate(&f, Favorite{
		UserID:    u.ID,
		ArticleId: a.ID,
	}).Error
	return err
}

func (a *Article) UnFavoriteBy(u *accounts.User) error {
	db := conf.GetDB()
	err := db.Where(Favorite{
		UserID:    u.ID,
		ArticleId: a.ID,
	}).Delete(Favorite{}).Error
	return err
}
