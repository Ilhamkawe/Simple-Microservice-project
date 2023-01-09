package orders

type CreateOrderInput struct {
	Status   string `json:"status" binding:"required"`
	UserID   int    `json:"user_id" binding:"required"`
	CourseID int    `json:"course_id" binding:"required"`
	Amount   int    `json:"amount" binding:"required"`
}

type MetadataInput struct {
	CourseID        int
	CoursePrice     int
	CourseName      string
	CourseThumbnail string
	CourseLevel     string
}
