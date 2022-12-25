package courses

type Service interface {
	Create(input CreateInputCourse) (Courses, error)
	FindCourseByID(id int) (Courses, error)
	Update(input UpdateInputCourse) (Courses, error)
	GetAll() ([]Courses, error)
	Destroy(id int) (bool, error)
	FilterCourse(name string, type_ string) ([]Courses, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Destroy(id int) (bool, error) {
	_, err := s.repository.Destroy(id)

	if err != nil {
		return false, nil
	}

	return true, nil
}

func (s *service) GetAll() ([]Courses, error) {
	var course []Courses
	course, err := s.repository.GetAll()

	if err != nil {
		return course, err
	}

	return course, nil
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

func (s *service) FindCourseByID(id int) (Courses, error) {
	var course Courses
	course, err := s.repository.FindCourseByID(id)

	if err != nil {
		return course, err
	}

	return course, nil

}

func (s *service) Update(input UpdateInputCourse) (Courses, error) {
	course := Courses{}
	course.ID = input.ID
	course.Name = input.Name
	course.Certificate = input.Certificate
	course.Thumbnail = input.Thumbnail
	course.Type = input.Type
	course.Status = input.Status
	course.Price = input.Price
	course.Level = input.Level
	course.Description = input.Description
	course.MentorID = input.MentorID

	newCourse, err := s.repository.Update(course)
	if err != nil {
		return newCourse, err
	}

	return newCourse, nil
}

func (s *service) FilterCourse(name string, type_ string) ([]Courses, error) {
	var courses []Courses
	courses, err := s.repository.FilterCourse(name, type_)
	if err != nil {
		return courses, err
	}

	return courses, nil
}

// ID          int
// Name        string
// Certificate bool
// Thumbnail   string
// Type        string
// Status      string
// Price       int
// Level       string
// Description string
// MentorID    int