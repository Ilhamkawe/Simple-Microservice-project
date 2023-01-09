package logs

type Service interface{}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}