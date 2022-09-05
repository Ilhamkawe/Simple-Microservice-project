package course

type Service interface {
	Create(input CreateInputMentor) (Mentors, error)
	IsEmailAvailable(email string) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
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
