package mentors

type Service interface {
	Create(input CreateInputMentor) (Mentors, error)
	IsEmailAvailable(email string) (bool, error)
	Update(input UpdateInputMentor) (Mentors, error)
	GetAllMentor() ([]Mentors, error)
	FindMentorByID(ID int) (Mentors, error)
	Destroy(id int) (bool, error)
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

func (s *service) FindMentorByID(id int) (Mentors, error) {
	var mentor Mentors
	mentor, err := s.repository.FindMentorByID(id)

	if err != nil {
		return mentor, err
	}
	return mentor, nil
}

// ambil semua data mentor
func (s *service) GetAllMentor() ([]Mentors, error) {
	mentors, err := s.repository.GetAllMentor()

	if err != nil {
		return mentors, err
	}

	return mentors, nil
}

// input data mentor
func (s *service) Create(input CreateInputMentor) (Mentors, error) {
	mentor := Mentors{}
	mentor.Name = input.Name
	mentor.Email = input.Email
	mentor.Profession = input.Profession
	mentor.Profile = input.Profile

	newMentor, err := s.repository.Create(mentor)

	if err != nil {
		return newMentor, err
	}

	return newMentor, nil

}

// cek apakah email tersedia
func (s *service) IsEmailAvailable(email string) (bool, error) {
	mentor, err := s.repository.FindMentorByEmail(email)

	if err != nil {
		return false, err
	}

	if mentor.ID == 0 {
		return true, nil
	}

	return false, nil
}

// update mentor
func (s *service) Update(input UpdateInputMentor) (Mentors, error) {
	mentor := Mentors{}
	mentor.Name = input.Name
	mentor.Email = input.Email
	mentor.Profile = input.Profile
	mentor.Profession = input.Profession
	newMentor, err := s.repository.Update(mentor)
	if err != nil {
		return newMentor, err
	}
	return newMentor, nil
}