package mycourse

type CreateInputMyCourse struct {
	CourseID int `json:"course_id" binding:"required"`
	UserID   int `json:"user_id" binding:"required"`
}