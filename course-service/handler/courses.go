package handler

import (
	"course-service/courses"
	"course-service/helper"
	"course-service/mentors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type coursesHandler struct {
	courseService courses.Service
	mentorsService mentors.Service
}

func NewCourseHandler(courseService courses.Service, mentorsService mentors.Service) *coursesHandler {
	return &coursesHandler{courseService, mentorsService}
}

func (h *coursesHandler) Create(c *gin.Context) {
	var input courses.CreateInputCourse
	err := c.ShouldBind(&input)

	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	mentor, err := h.mentorsService.FindMentorByID(input.MentorID)

	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if mentor.ID == 0 {
		response := helper.ApiResponse("Mentor doesnt exist", http.StatusBadRequest, "error", gin.H{
			
		})

		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	course, err := h.courseService.Create(input)

	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Input Course", http.StatusOK, "Success", courses.FormatCourses(course))
	c.JSON(http.StatusOK, response)
}