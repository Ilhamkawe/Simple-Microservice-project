package mentors

type CreateInputMentor struct {
	Name       string `json:"name" binding:"required"`
	Profile    string `json:"profile" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Profession string `json:"profession" binding:"required"`
}
type UpdateInputMentor struct {
	Name       string `json:"name" binding:"required"`
	Profile    string `json:"profile" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Profession string `json:"profession" binding:"required"`
}

type UpdateMentorURI struct {
	ID int `uri:"id"`
}

type InputUriID struct {
	ID int `uri:"id"`
}
