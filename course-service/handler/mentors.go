package handler

import (
	"course-service/helper"
	"course-service/mentors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type mentorsHandler struct {
	mentorsService mentors.Service
}

func NewMentorHandler(mentorsService mentors.Service) *mentorsHandler {
	return &mentorsHandler{mentorsService}
}

// !======================================================

func (h *mentorsHandler) Destroy(c *gin.Context){

	var uri mentors.InputUriID
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

	mentor, err := h.mentorsService.Destroy(uri.ID)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Delete Data", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Delete Data", http.StatusOK, "success", mentor)
	c.JSON(http.StatusOK, response)

}

// !======================================================

func (h *mentorsHandler) Index(c *gin.Context){

	mentor, err := h.mentorsService.GetAllMentor()
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error when fetching data", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("success get data", http.StatusOK, "success", mentors.FormatMentors(mentor))
	c.JSON(http.StatusOK, response)

	
}

// !======================================================

func (h *mentorsHandler) Show(c *gin.Context){
	var uri mentors.InputUriID
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

	println(uri.ID)

	mentor, err := h.mentorsService.FindMentorByID(uri.ID)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Check Email", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("sucess get data", http.StatusOK, "success", mentors.FormatCourse(mentor))
	c.JSON(http.StatusOK, response)

}

// !======================================================

func (h *mentorsHandler) UpdateMentors(c *gin.Context) {
	var uri mentors.UpdateMentorURI
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

	var input mentors.UpdateInputMentor 
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

	mentor, err := h.mentorsService.FindMentorByID(uri.ID)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Check Email", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	emailAvailable, err := h.mentorsService.IsEmailAvailable(input.Email)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Check Email", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	if mentor.Email != input.Email && !emailAvailable {
		response := helper.ApiResponse("Email Already Exist", http.StatusConflict, "error", gin.H{})

		c.JSON(http.StatusConflict, response)
		return
	}

	newMentor, err := h.mentorsService.Update(input)

	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Update Mentors", http.StatusOK, "Success", mentors.FormatCourse(newMentor))
	c.JSON(http.StatusOK, response)

}

// !======================================================

func (h *mentorsHandler) CreateMentors(c *gin.Context) {
	var input mentors.CreateInputMentor
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

	emailAvailable, err := h.mentorsService.IsEmailAvailable(input.Email)
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

	newMentor, err := h.mentorsService.Create(input)

	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Input Mentors", http.StatusOK, "Success", mentors.FormatCourse(newMentor))
	c.JSON(http.StatusOK, response)
}
