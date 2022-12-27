package handler

import (
	"course-service/courses"
	"course-service/helper"
	mycourse "course-service/mycourses"
	"net/http"
	"strconv"

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
		response := helper.ApiResponse("Error When Check Course", http.StatusBadRequest, "error", errorMessage)

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
		response := helper.ApiResponse("Error When Input Data", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success Input ", http.StatusOK, "Success", mycourse.FormatMyCourse(newMycourse))
	c.JSON(http.StatusOK, response)
}

func (h *mycoursesHandler) Index(c *gin.Context){
	user_id := c.Query("user_id")


	if user_id != "" {
		
		u_id, err := strconv.Atoi(user_id)
		if err != nil {
			error := helper.FormatValidationError(err)
			errorMessage := gin.H{
				"error": error,
			}
			response := helper.ApiResponse("Error when Getting Params", http.StatusBadRequest, "error", errorMessage)
	
			c.JSON(http.StatusBadRequest, response)
			return
		}

		myCourse, err := h.MyCoursesService.FindByUserID(u_id)

		if err != nil {
			
			response := helper.ApiResponse("Error when fetching data", http.StatusBadRequest, "error", err)
	
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.ApiResponse("success get data", http.StatusOK, "success", myCourse)
		c.JSON(http.StatusOK, response)
		return
	}

	myCourse, err := h.MyCoursesService.GetAll()
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": error,
		}
		response := helper.ApiResponse("Error when fetching data", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("success get data", http.StatusOK, "success", mycourse.FormatMyCourses(myCourse))
	c.JSON(http.StatusOK, response)
	
}