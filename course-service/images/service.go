package images

type Service interface {
	Create(input CreateInputImage) (ImageCourses, error)
	Destroy(id int) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input CreateInputImage) (ImageCourses, error) {
	img := ImageCourses{}
	img.CourseID = input.CourseID
	img.Image = input.Image

	newImg, err := s.repository.Create(img)
	if err != nil {
		return newImg, err
	}

	return newImg, nil
}

func (s *service) Destroy(id int) (bool, error) {
	img, err := s.repository.Destroy(id)
	if err != nil {
		return img, err
	}

	return img, nil
}