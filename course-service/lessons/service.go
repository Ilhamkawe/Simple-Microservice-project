package lessons

type Service interface {
	GetAll() ([]Lessons, error)
	FindByCHapterID(id int) ([]Lessons, error)
	Create(input CreateInputLessons) (Lessons, error)
	Destroy(id int) (bool, error)
	Update(input UpdateInputLessons) (Lessons, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input CreateInputLessons) (Lessons, error) {
	lesson := Lessons{}
	lesson.Name = input.Name
	lesson.Video = input.Video
	lesson.ChapterID = input.ChapterID

	newLesson, err := s.repository.Create(lesson)
	if err != nil {
		return newLesson, err
	}

	return newLesson, nil
}

func (s *service) Update(input UpdateInputLessons) (Lessons, error) {
	lesson := Lessons{}
	lesson.ID = input.ID
	lesson.Name = input.Name
	lesson.Video = input.Video
	lesson.ChapterID = input.ChapterID

	newLesson, err := s.repository.Update(lesson)
	if err != nil {
		return newLesson, err
	}

	return newLesson, nil
}

func (s *service) FindByID(id int) (Lessons, error) {
	var lesson Lessons

	lesson, err := s.repository.FindByID(id)

	if err != nil {
		return lesson, err
	}

	return lesson, nil
}

func (s *service) FindByCHapterID(id int) ([]Lessons, error) {
	var lessons []Lessons

	lessons, err := s.repository.FindByChapterID(id)

	if err != nil {
		return lessons, err
	}

	return lessons, nil
}

func (s *service) GetAll() ([]Lessons, error) {
	var lessons []Lessons

	lessons, err := s.repository.GetAll()

	if err != nil {
		return lessons, err
	}

	return lessons, nil
}

func (s *service) Destroy(id int) (bool, error) {
	lessons, err := s.repository.Destroy(id)

	if err != nil {
		return lessons, err
	}

	return lessons, nil
}
