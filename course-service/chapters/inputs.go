package chapters

type CreateInputChapters struct {
	Name     string `json:"name" binding:"required"`
	CourseID int    `json:"course_id" binding:"required"`
}
type UpdateInputChapters struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	CourseID int    `json:"course_id" binding:"required"`
}
type InputID struct {
	ID int `json:"id" binding:"required"`
}
type UriID struct {
	ID int `uri:"id" binding:"required"`
}