package images

type CreateInputImage struct {
	CourseID int    `json:"course_id" binding:"required"`
	Image    string `json:"image" binding:"required"`
}

type InputUriID struct {
	ID int `uri:"id"`
}