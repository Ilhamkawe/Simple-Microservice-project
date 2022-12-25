package handler

import (
	"course-service/courses"
	"course-service/helper"
	mycourse "course-service/mycourses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type mycoursesHandler struct {
	MyCoursesService mycourse.Service
	CoursesService courses.Service
}

func NewMyCoursesHandler(mycoursesService mycourse.Service, CoursesService courses.Service) *mycoursesHandler {
	return &mycoursesHandler{mycoursesService, CoursesService}
}

func (h *mycoursesHandler) Create(c *gin.Context) {

	var input mycourse.CreateInputMyCourse 
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

	courseAvailable, err := h.MyCoursesService.IsEnrolled(input)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Check Email", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !courseAvailable  {
		response := helper.ApiResponse("Already Enrol", http.StatusConflict, "error", gin.H{})

		c.JSON(http.StatusConflict, response)
		return
	}

	newMycourse, err := h.MyCoursesService.Create(input)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Check Email", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Input ", http.StatusOK, "Success", newMycourse)
	c.JSON(http.StatusOK, response)


}