package common

import (
	"github.com/rezaAmiri123/library/conf"
)

func FindObject(object, condition interface{}) error {
	db := conf.GetDB()
	return db.Where(condition).First(object).Error
}

func SaveObject(data interface{}) error {
	db := conf.GetDB()
	return db.Save(data).Error
}