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


// !======================================================
func (h *coursesHandler) Destroy(c *gin.Context){

	var uri courses.InputUriID
	err := c.ShouldBindUri(&uri)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid Uri", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	course, err := h.courseService.Destroy(uri.ID)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Delete Data", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Delete Data", http.StatusOK, "success", course)
	c.JSON(http.StatusOK, response)

}

// !======================================================
func (h *coursesHandler) Index(c *gin.Context){
	name := c.Query("q")
	type_ := c.Query("type")

	// filtering
	if name != "" || type_ != "" {
		course, err := h.courseService.FilterCourse(name, type_)
		if err != nil {
			
			response := helper.ApiResponse("Error when fetching data", http.StatusBadRequest, "error", err)
	
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.ApiResponse("success get data", http.StatusOK, "success", courses.FormatCourses(course))
		c.JSON(http.StatusOK, response)
		return

	}

	course, err := h.courseService.GetAll()
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error when fetching data", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("success get data", http.StatusOK, "success", courses.FormatCourses(course))
	c.JSON(http.StatusOK, response)
}

// !======================================================
func (h *coursesHandler) Update(c *gin.Context){
	var uri courses.InputUriID 
	err := c.ShouldBindUri(&uri)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid Uri", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input courses.UpdateInputCourse 
	input.ID = uri.ID
	err = c.ShouldBind(&input)

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
		response := helper.ApiResponse("Mentor doesnt exist", http.StatusBadRequest, "error", gin.H{})

		c.JSON(http.StatusBadRequest, response)
		return
	}

	newCourse, err := h.courseService.Update(input)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Update Courses", http.StatusOK, "Success", courses.FormatCourse(newCourse))
	c.JSON(http.StatusOK, response)
}

// !======================================================
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

	response := helper.ApiResponse("Success Input Course", http.StatusOK, "Success", courses.FormatCourse(course))
	c.JSON(http.StatusOK, response)
}