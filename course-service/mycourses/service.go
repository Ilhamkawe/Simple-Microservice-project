package mycourse

type Service interface {
	FindByUserID(id int) ([]MyCourses, error)
	GetAll() ([]MyCourses, error)
	Create(input CreateInputMyCourse) (MyCourses, error)
	IsEnrolled(input CreateInputMyCourse) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindByUserID(id int) ([]MyCourses, error) {
	var mycourse []MyCourses

	mycourse, err := s.repository.FindByUserID(id)

	if err != nil {
		return mycourse, err
	}

	return mycourse, nil
}

func (s *service) GetAll() ([]MyCourses, error) {

	var mycourse []MyCourses

	mycourse, err := s.repository.GetAll()

	if err != nil {
		return mycourse, err
	}

	return mycourse, nil

}

func (s *service) Create(input CreateInputMyCourse) (MyCourses, error) {
	mycourse := MyCourses{}
	mycourse.CourseID = input.CourseID
	mycourse.UserID = input.UserID

	newmycourse, err := s.repository.Create(mycourse)
	if err != nil {
		return newmycourse, err
	}

	return newmycourse, nil

}

func (s *service) IsEnrolled(input CreateInputMyCourse) (bool, error) {

	mycourse, err := s.repository.FindByUserAndCourseID(input.UserID, input.CourseID)

	if err != nil {
		return false, err
	}

	if mycourse.ID == 0 {
		return true, nil
	}

	return false, nil

}