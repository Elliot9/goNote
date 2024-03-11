package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database struct
type Database struct {
	DB *gorm.DB
}

func Connect(user, password, host, dbName string) Database {
	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s", user, password,
		host, dbName)
	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		panic("Failed to connect to database!")
	}

	return Database{
		DB: db,
	}
}
