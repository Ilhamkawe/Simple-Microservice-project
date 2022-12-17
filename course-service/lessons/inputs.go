package lessons

type CreateInputLessons struct {
	Name      string `json:"name" binding:"required"`
	Video     string `json:"video" binding:"required"`
	ChapterID int    `json:"chapter_id" binding:"required"`
}

type UpdateInputLessons struct {
	ID        int    `json:"id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Video     string `json:"video" binding:"required"`
	ChapterID int    `json:"chapter_id" binding:"required"`
}

type InputUriID struct {
	ID int `uri:"id"`
}