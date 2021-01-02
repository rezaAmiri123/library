package common

import (
	"github.com/rezaAmiri123/library/conf"
)

func FindObject(condition interface{}, model *interface{}) error {
	db := conf.GetDB()
	return db.Where(condition).First(model).Error
}
