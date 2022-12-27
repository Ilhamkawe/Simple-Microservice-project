package handler

import (
	"course-service/courses"
	"course-service/helper"
	"course-service/riviews"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RiviewsHandler struct {
	RiviewService riviews.Service
	CourseService courses.Service
}

func NewRiviewsHandler(riviewsService riviews.Service, coursesService courses.Service) *RiviewsHandler {
	return &RiviewsHandler{riviewsService, coursesService}
}

func (h *RiviewsHandler) Create(c *gin.Context) {
	var input riviews.CreateInputRiview
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

	course, err := h.CourseService.FindCourseByID(input.CourseID)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Trying Check Course", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	if course.ID == 0 {
		response := helper.ApiResponse("Already Enrol", http.StatusConflict, "error", gin.H{})

		c.JSON(http.StatusConflict, response)
		return
	}

	isRiviewAvailable, err := h.RiviewService.IsRiviewAvailable(input.UserID, input.CourseID)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Trying Check Riview", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !isRiviewAvailable  {
		response := helper.ApiResponse("Already Riviewed", http.StatusConflict, "error", gin.H{})

		c.JSON(http.StatusConflict, response)
		return
	}

	newRiview, err := h.RiviewService.Create(input)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Check Email", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Input ", http.StatusOK, "Success", riviews.FormatRiview(newRiview))
	c.JSON(http.StatusOK, response)

}

func (h *RiviewsHandler) Destroy(c *gin.Context) {
	var uri riviews.UriID

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

	riview, err := h.RiviewService.Destroy(uri.ID)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Delete Data", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Delete Data", http.StatusOK, "success", riview)
	c.JSON(http.StatusOK, response)
}

func (h *RiviewsHandler) Update(c *gin.Context) {
	var uri riviews.UriID

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

	var input riviews.UpdateInputRiview 
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

	riview, err := h.RiviewService.GetByID(uri.ID) 

	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Fetching Data", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	if riview.ID == 0 {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Riview Not Found", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	newRiview, err := h.RiviewService.Update(input)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error WHen Update Data", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Delete Data", http.StatusOK, "success", riviews.FormatRiview(newRiview))
	c.JSON(http.StatusOK, response)

}