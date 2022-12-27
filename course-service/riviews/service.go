package riviews

type Service interface {
	Create(input CreateInputRiview) (Riviews, error)
	IsRiviewAvailable(user_id int, course_id int) (bool, error)
	Update(input UpdateInputRiview) (Riviews, error)
	Destroy(id int) (bool, error)
	GetByID(id int) (Riviews, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// menginput data riview pada course
func (s *service) Create(input CreateInputRiview) (Riviews, error) {
	riview := Riviews{}
	riview.UserID = input.UserID
	riview.CourseID = input.CourseID
	riview.Rating = input.Rating
	riview.Note = input.Note

	newRiview, err := s.repository.Create(riview)

	if err != nil {
		return newRiview, err
	}

	return newRiview, nil
}

// apakah user sudah memberikan riview pada course
func (s *service) IsRiviewAvailable(user_id int, course_id int) (bool, error) {
	riview, err := s.repository.FindByUserAndCourseID(user_id, course_id)

	if err != nil {
		return false, nil
	}

	if riview.ID == 0 {
		return true, nil
	}

	return false, nil
}

// mengupdate data riview pada course
func (s *service) Update(input UpdateInputRiview) (Riviews, error) {
	riview := Riviews{}

	riview.Rating = input.Rating
	riview.Note = input.Note

	newRiview, err := s.repository.Update(riview)

	if err != nil {
		return newRiview, err
	}

	return newRiview, nil
}

// menghapus data riview pada course
func (s *service) Destroy(id int) (bool, error) {
	_, err := s.repository.Destroy(id)

	if err != nil {
		return false, nil
	}

	return true, nil
}

// mengambil data riview
func (s *service) GetByID(id int) (Riviews, error) {
	riview, err := s.repository.GetByID(id)

	if err != nil {
		return riview, err
	}

	return riview, nil
}