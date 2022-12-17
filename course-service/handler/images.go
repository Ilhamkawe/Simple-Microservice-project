package handler

import (
	"course-service/courses"
	"course-service/helper"
	"course-service/images"
	"encoding/base64"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ImagesHandler struct {
	courseService courses.Service
	ImagesService images.Service
}

func NewImageHandler(courseService courses.Service,imagesService images.Service) *ImagesHandler {
	return &ImagesHandler{courseService, imagesService}
}

// !======================================================

func (h *ImagesHandler) Create(c *gin.Context){
	var input images.CreateInputImage
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

	course, err := h.courseService.FindCourseByID(input.CourseID)
	
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

	i:= strings.Index(input.Image, ",")

	if i < 0 {
		log.Fatal("no comma")
	}
	data, _ := base64.StdEncoding.DecodeString(input.Image[i+1:])
	r := string(data)
	// get max mime index 
	coI := strings.Index(string(input.Image), ",")
	ext := strings.TrimSuffix(input.Image[11:coI], ";base64")

	filename := "images/gallery/"+"course-img-"+strconv.Itoa(course.ID)+"-"+strconv.Itoa(rand.Intn(9999))+"."+ext

	err = ioutil.WriteFile(filename, []byte(r), 0644)
	if err != nil {
		response := helper.ApiResponse("Error while write image", http.StatusBadRequest, "error", gin.H{
			
		})

		c.JSON(http.StatusBadRequest, response)
		return
	}

	input.Image = filename
	
	newImg, err := h.ImagesService.Create(input)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Invalid input", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Input Image", http.StatusOK, "Success", images.FormatImageCourse(newImg))
	c.JSON(http.StatusOK, response)


}

// !======================================================

func (h *ImagesHandler) Destroy(c *gin.Context) {

	var uri images.InputUriID
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

	image, err := h.ImagesService.Destroy(uri.ID)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error When Delete Data", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Delete Data", http.StatusOK, "success", image)
	c.JSON(http.StatusOK, response)
}

