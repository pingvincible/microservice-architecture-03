package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	//dsn := "host=arch.homework user=cruduser password=crudpass dbname=crudapp port=30856 sslmode=disable TimeZone=Asia/Almaty"
	dsn := os.Getenv("DATABASE_URI")
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
