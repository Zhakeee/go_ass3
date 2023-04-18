package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	dsn := "host=localhost user=postgres password=12345 dbname=a sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("No connection with db")
	} else {
		fmt.Println("Successful connection with db")
	}
	return db
}