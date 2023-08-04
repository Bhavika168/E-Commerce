package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitialiseDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=123456 dbname=first_db port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// fmt.Println(dsn)
	return db
}
