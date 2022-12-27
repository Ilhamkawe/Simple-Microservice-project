package lessons

import (
	"time"
)

type Lessons struct {
	ID        int
	Name      string
	Video     string
	ChapterID int
	CreatedAt time.Time
	UpdatedAt time.Time
}