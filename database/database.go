package database

import (
	"dumbmerch-api/models/entity"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	dsn := "root:@tcp(127.0.0.1:3306)/dumbmerch-api?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("cannot connect database")
	}

	fmt.Println("connected database")
}

func Migration() {
	err := DB.AutoMigrate(
		&entity.User{},
		&entity.Profile{},
		&entity.Product{},
		&entity.Transaction{},
	)

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
