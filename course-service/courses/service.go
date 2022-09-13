package courses

type Service interface {
	Create(input CreateInputCourse) (Courses, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input CreateInputCourse) (Courses, error) {
	course := Courses{}
	course.Name = input.Name
	course.Certificate = input.Certificate
	course.Description = input.Description
	course.Level = input.Level
	course.Price = input.Price
	course.Thumbnail = input.Thumbnail
	course.Type = input.Type
	course.Status = input.Status
	course.MentorID = input.MentorID

	newCourse, err := s.repository.Create(course)

	if err != nil {
		return newCourse, err
	}

	return newCourse, nil

}