package course

type CreateInputMentor struct {
	Name       string `json:"name" binding:"required"`
	Profile    string `json:"profile" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Profession string `json:"profession" binding:"required"`
}
