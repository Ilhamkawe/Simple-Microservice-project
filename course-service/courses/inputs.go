package courses

type CreateInputCourse struct {
	Name        string `json:"name" binding:"required"`
	Certificate bool   `json:"certificate" binding:"required"`
	Thumbnail   string `json:"thumbnail" binding:"required"`
	Price       int    `json:"price"`
	Status      string `json:"status" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Level       string `json:"level" binding:"required"`
	Description string `json:"description" binding:"required"`
	MentorID    int    `json:"mentor_id" binding:"required"`
}

// name
// certificate
// thumbnail
// price
// status
// type
// level
// description
// mentor_id