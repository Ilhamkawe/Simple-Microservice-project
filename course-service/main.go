package main

import (
	"course-service/chapters"
	"course-service/courses"
	"course-service/handler"
	"course-service/images"
	"course-service/lessons"
	"course-service/mentors"
	mycourse "course-service/mycourses"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/service_course?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// * Course Dependencies
	mentorRepository := mentors.NewRepository(db)
	mentorService := mentors.NewService(mentorRepository)
	mentorHandler := handler.NewMentorHandler(mentorService)

	courseRepository := courses.NewRepository(db)
	courseService := courses.NewService(courseRepository)
	courseHandler := handler.NewCourseHandler(courseService, mentorService)

	chapterRepository := chapters.NewRepository(db)
	chapterService := chapters.NewService(chapterRepository)
	chapterHandler := handler.NewChapterHandler(chapterService, courseService)

	lessonRepository := lessons.NewRepository(db)
	lessonService := lessons.NewService(lessonRepository)
	lessonHandler := handler.NewLessonHandler(lessonService, chapterService)

	imageRepository := images.NewRepository(db)
	imageService := images.NewService(imageRepository)
	imageHandler := handler.NewImageHandler(courseService,imageService) 

	myCourseRepository := mycourse.NewRepository(db)
	myCourseService := mycourse.NewService(myCourseRepository)
	myCourseHandler := handler.NewMyCoursesHandler(myCourseService, courseService)

	router := gin.Default()

	api := router.Group("/api/v1")

	// mentors
	api.POST("/mentors", mentorHandler.CreateMentors)
	api.PUT("/mentors/:id", mentorHandler.UpdateMentors)
	api.GET("/mentors", mentorHandler.Index)
	api.GET("/mentors/:id", mentorHandler.Show)
	api.DELETE("/mentors/:id", mentorHandler.Destroy)

	// courses
	api.POST("/courses", courseHandler.Create)
	api.PUT("/courses/:id", courseHandler.Update)
	api.GET("/courses", courseHandler.Index)
	api.DELETE("/courses/:id", courseHandler.Destroy)

	// chapters
	api.POST("/chapters", chapterHandler.Create)
	api.GET("/chapters", chapterHandler.GetChapters)
	api.GET("/chapters/:id", chapterHandler.GetDetail)
	api.DELETE("/chapters/:id", chapterHandler.Destroy)
	api.PUT("/chapters/:id", chapterHandler.Update)

	// lessons
	api.POST("/lessons", lessonHandler.Create)
	api.PUT("/lessons/:id", lessonHandler.Update)
	api.GET("/lessons/:id", lessonHandler.Index)
	api.DELETE("/lessons/:id", lessonHandler.Destroy)

	// course image
	api.POST("/course/image", imageHandler.Create)
	api.DELETE("/course/image/:id", imageHandler.Destroy)

	// mycourses
	api.GET("/mycourses", myCourseHandler.Index)
	api.POST("/mycourses", myCourseHandler.Create)

	// riviewss

	router.Run(":3002")

}
