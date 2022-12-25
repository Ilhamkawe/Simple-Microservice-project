package handler

import (
	"course-service/chapters"
	"course-service/courses"
	"course-service/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type chapterHandler struct {
	chapterService chapters.Service
	CoursesService courses.Service
}

func NewChapterHandler(chapterService chapters.Service, courseService courses.Service) *chapterHandler {
	return &chapterHandler{chapterService, courseService}
}

// !======================================================
func (h *chapterHandler) Create(c *gin.Context){
	var input chapters.CreateInputChapters
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

	course, err := h.CoursesService.FindCourseByID(input.CourseID)

	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if course.ID == 0 {
		response := helper.ApiResponse("Course doesnt exist", http.StatusBadRequest, "error", gin.H{
			
		})

		c.JSON(http.StatusBadRequest, response)
		return
	}

	chapter, err := h.chapterService.Create(input)

	if err != nil { 
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Input Chapter", http.StatusOK, "Success", chapters.FormatShowChapters(chapter))
	c.JSON(http.StatusOK, response)
}
// !======================================================
func (h *chapterHandler) Update(c *gin.Context) {
	var uri chapters.UriID
	err:= c.ShouldBindUri(&uri)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid Uri", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input chapters.UpdateInputChapters
	input.ID = uri.ID
	err = c.ShouldBindJSON(&input)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newChapter, err := h.chapterService.Update(input)
	if err != nil { 
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Update Chapter", http.StatusOK, "Success", chapters.FormatShowChapters(newChapter))
	c.JSON(http.StatusOK, response)

}
// !======================================================
func (h *chapterHandler) GetChapters(c *gin.Context){

	courseID := c.Query("course_id")
	id, _ := strconv.Atoi(courseID)

	if courseID == "" {
		Vchapters, err :=  h.chapterService.GetAllChapters()
		if err != nil {
			error := helper.FormatValidationError(err)
			errorMessage := gin.H{
				"error": error,
			}
			response := helper.ApiResponse("Something Wrong", http.StatusUnprocessableEntity, "error", errorMessage)

			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		response := helper.ApiResponse("Success Get Chapter", http.StatusOK, "Success", chapters.FormatAllChapters(Vchapters))
		c.JSON(http.StatusOK, response)

		return
	}

	Vchapters, err := h.chapterService.GetAllCourseChapters(id)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Something Wrong", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Success Get Chapter", http.StatusOK, "Success", chapters.FormatAllChapters(Vchapters))
	c.JSON(http.StatusOK, response)
}
// !======================================================
func (h* chapterHandler) GetDetail(c *gin.Context){
	var input chapters.UriID
	err := c.ShouldBindUri(&input)

	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	chapter, err := h.chapterService.Detail(input.ID)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Something Wrong", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	fmt.Println(chapter)

	response := helper.ApiResponse("Success Get Chapter", http.StatusOK, "Success", chapters.FormatShowChapters(chapter))
	c.JSON(http.StatusOK, response)
}
// !======================================================
func (h* chapterHandler) Destroy(c *gin.Context){
	var input chapters.UriID
	err := c.ShouldBindUri(&input)

	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	chapter, err := h.chapterService.Destroy(input.ID)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Delete Data", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Delete Data", http.StatusOK, "success", chapter)
	c.JSON(http.StatusOK, response)

}
