package handler

import (
	"course-service/chapters"
	"course-service/helper"
	"course-service/lessons"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type lessonsHandler struct {
	lessonService lessons.Service
	chapterService chapters.Service
}

func NewLessonHandler(lessonService lessons.Service, chapterService chapters.Service) *lessonsHandler{
	return &lessonsHandler{lessonService, chapterService}
}

// !======================================================
// get all lesson list or all lesson by chapter id using query chapter_id
func (h* lessonsHandler) Index(c *gin.Context){

	chapter_id := c.Query("chapter_id")

	if chapter_id != "" {
		ch_id, _ := strconv.Atoi(chapter_id)
		lessons, err := h.lessonService.FindByCHapterID(ch_id)
		if err != nil {
			error := helper.FormatValidationError(err)
			errorMessage := gin.H{
				"error": error,
			}
			response := helper.ApiResponse("Error when fetching data", http.StatusBadRequest, "error", errorMessage)

			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.ApiResponse("success get data", http.StatusOK, "success", lessons)
		c.JSON(http.StatusOK, response)
		return

	}

	allLessons, err := h.lessonService.GetAll()
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error when fetching data", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("success get data", http.StatusOK, "success", lessons.FormatLessons(allLessons))
	c.JSON(http.StatusOK, response)

}

// !======================================================

func (h *lessonsHandler) Create(c *gin.Context){

	var input lessons.CreateInputLessons
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

	ch, err := h.chapterService.Detail(input.ChapterID)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if ch.ID == 0 {
		response := helper.ApiResponse("Chapter doesnt exist", http.StatusBadRequest, "error", gin.H{
			
		})

		c.JSON(http.StatusBadRequest, response)
		return
	}

	newLessons, err := h.lessonService.Create(input)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Input Course", http.StatusOK, "Success", lessons.FormatLesson(newLessons))
	c.JSON(http.StatusOK, response)
}

// !======================================================

func (h *lessonsHandler) Destroy(c *gin.Context){
	var uri lessons.InputUriID
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

	lessons, err := h.lessonService.Destroy(uri.ID)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Delete Data", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Delete Data", http.StatusOK, "success", lessons)
	c.JSON(http.StatusOK, response)
}

// !======================================================

func (h *lessonsHandler) Update(c *gin.Context){
	var uri lessons.InputUriID
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

	var input lessons.UpdateInputLessons
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

	ch, err := h.chapterService.Detail(input.ChapterID)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if ch.ID == 0 {
		response := helper.ApiResponse("Chapter doesnt exist", http.StatusBadRequest, "error", gin.H{
			
		})

		c.JSON(http.StatusBadRequest, response)
		return
	}

	newLesson, err := h.lessonService.Update(input)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Update Lesson", http.StatusOK, "Success", lessons.FormatLesson(newLesson))
	c.JSON(http.StatusOK, response)
}

// create //done
// update //done
// get list & detail //done
// delete //done