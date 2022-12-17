package lessons

import (
	"course-service/chapters"
	"time"
)

type Lessons struct {
	ID        int
	Name      string
	Video     string
	ChapterID int
	Chapter chapters.Chapters
	CreatedAt time.Time
	UpdatedAt time.Time
}