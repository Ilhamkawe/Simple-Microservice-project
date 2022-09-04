package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// * Course Dependencies
	courseRepository = 

}
