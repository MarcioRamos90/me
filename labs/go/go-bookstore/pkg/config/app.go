package config

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(
		"mysql",
		"root:#10293848FgT@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	d.DB().SetConnMaxLifetime(time.Second)
	db = d
}

func GetDB() *gorm.DB {
	return db
}
