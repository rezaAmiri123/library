package common

import (
	"github.com/rezaAmiri123/library/conf"
)

func FindObject(model *interface{}, condition interface{}) error {
	db := conf.GetDB()
	return db.Where(condition).First(model).Error
}

func SaveObject(data interface{}) error {
	db := conf.GetDB()
	return db.Create(data).Error
}