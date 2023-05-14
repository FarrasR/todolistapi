package database

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB
var once sync.Once

func InitDB() *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open(mysql.Open(getDSN()))

		if err != nil {
			panic(err)
		}
		dbInstance = db
	})
	return dbInstance
}

func getDSN() string {
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"))
	return DSN
}
