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
	Chapters []Chapters `gorm:"foreignKey:CourseID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Chapters struct {
	ID       int 
	Name     string 
	CourseID int 
	Lessons []Lessons `gorm:"foreignKey:ChapterID"`
	CreatedAt time.Time 
	UpdatedAt time.Time  
}

type Lessons struct {
	ID        int
	Name      string
	Video     string
	ChapterID int
	CreatedAt time.Time
	UpdatedAt time.Time
}