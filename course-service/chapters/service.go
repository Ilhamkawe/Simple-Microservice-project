package chapters

type Service interface {
	Create(input CreateInputChapters) (Chapters, error)
	GetAllChapters() ([]Chapters, error)
	GetAllCourseChapters(id int) ([]Chapters, error)
	Detail(id int) (Chapters, error)
	Destroy(id int) (bool, error)
	Update(input UpdateInputChapters) (Chapters, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllChapters() ([]Chapters, error) {
	var chapters []Chapters

	chapters, err := s.repository.Index()

	if err != nil {
		return chapters, err
	}

	return chapters, nil
}

func (s *service) GetAllCourseChapters(id int) ([]Chapters, error) {
	var chapters []Chapters

	chapters, err := s.repository.FindByCourseID(id)

	if err != nil {
		return chapters, nil
	}

	return chapters, err
}

func (s *service) Create(input CreateInputChapters) (Chapters, error) {
	chapter := Chapters{}
	chapter.Name = input.Name
	chapter.CourseID = input.CourseID

	newChapters, err := s.repository.Create(chapter)

	if err != nil {
		return newChapters, err
	}

	return newChapters, nil
}

func (s *service) Update(input UpdateInputChapters) (Chapters, error) {
	chapter := Chapters{}
	chapter.ID = input.ID
	chapter.Name = input.Name
	chapter.CourseID = input.CourseID

	newChapters, err := s.repository.Update(chapter)

	if err != nil {
		return newChapters, err
	}

	return newChapters, nil
}

func (s *service) Detail(id int) (Chapters, error) {
	var chapter Chapters

	chapter, err := s.repository.FindByID(id)

	if err != nil {
		return chapter, err
	}

	return chapter, nil
}

func (s *service) Destroy(id int) (bool, error) {
	_, err := s.repository.Destroy(id)

	if err != nil {
		return false, nil
	}

	return true, nil
}
