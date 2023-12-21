package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// connect to database
func Connect() {
	// user name; password; name of the database; ? ...
	// ?charset=utf8&parseTime=True&loc=Local
	d, err := gorm.Open("mysql", "root:000609/learngo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
