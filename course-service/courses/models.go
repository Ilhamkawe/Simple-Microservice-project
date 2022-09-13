package courses

import (
	"course-service/mentors"
	"time"
)

type Courses struct {
	ID          int
	Name        string
	Certificate bool
	Thumbnail   string
	Type        string
	Status      string
	Price       int
	Level       string
	Description string
	MentorID    int
	Mentor      mentors.Mentors
	CreatedAt   time.Time
	UpdatedAt   time.Time
}