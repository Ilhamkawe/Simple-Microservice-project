package main

import (
	"course-service/courses"
	"course-service/handler"
	"course-service/mentors"
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
	mentorRepository := mentors.NewRepository(db)
	mentorService := mentors.NewService(mentorRepository)
	mentorHandler := handler.NewMentorHandler(mentorService)

	courseRepository := courses.NewRepository(db)
	courseService := courses.NewService(courseRepository)
	courseHandler := handler.NewCourseHandler(courseService, mentorService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/mentors", mentorHandler.CreateMentors)
	api.PUT("/mentors/:id", mentorHandler.UpdateMentors)
	api.GET("/mentors", mentorHandler.Index)
	api.GET("/mentors/:id", mentorHandler.Show)
	api.DELETE("/mentors/:id", mentorHandler.Destroy)

	api.POST("/courses", courseHandler.Create)
	api.PUT("/courses/:id", courseHandler.Update)
	api.GET("/courses", courseHandler.Index)
	api.DELETE("/courses/:id", courseHandler.Destroy)

	router.Run(":3002")

}