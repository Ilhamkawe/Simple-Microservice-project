package handler

import (
	"course-service/course"
	"course-service/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type courseHandler struct {
	courseService course.Service
}

func NewCourseHandler(courseService course.Service) *courseHandler {
	return &courseHandler{courseService}
}

func (h *courseHandler) UpdateMentors(c *gin.Context) {

}

func (h *courseHandler) CreateMentors(c *gin.Context) {
	var input course.CreateInputMentor
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

	emailAvailable, err := h.courseService.IsEmailAvailable(input.Email)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Check Email", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !emailAvailable {
		response := helper.ApiResponse("Email Already Exist", http.StatusConflict, "error", gin.H{})

		c.JSON(http.StatusConflict, response)
		return
	}

	newMentor, err := h.courseService.Create(input)

	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Input Mentors", http.StatusOK, "Success", course.FormatCourse(newMentor))
	c.JSON(http.StatusOK, response)
}
