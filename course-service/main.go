package main

import (
	"course-service/course"
	"course-service/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/service_course?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// * Course Dependencies
	courseRepository := course.NewRepository(db)
	courseService := course.NewService(courseRepository)
	courseHandler := handler.NewCourseHandler(courseService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/mentors", courseHandler.CreateMentors)
	router.Run(":3002")

}
