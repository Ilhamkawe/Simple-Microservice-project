package riviews

type CreateInputRiview struct {
	UserID   int    `json:"user_id" binding:"required"`
	CourseID int    `json:"course_id" binding:"required"`
	Rating   int    `json:"rating" binding:"required"`
	Note     string `json:"note" binding:"required"`
}

type UriID struct {
	ID int `uri:"id" binding:"required"`
}

type UpdateInputRiview struct {
	Rating int    `json:"rating" binding:"required"`
	Note   string `json:"note" binding:"required"`
}