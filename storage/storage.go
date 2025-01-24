package storage

import (
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() (*gorm.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?parseTime=True"
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(6)
	sqlDB.SetConnMaxLifetime(time.Minute * 1)

	return db, err
}
