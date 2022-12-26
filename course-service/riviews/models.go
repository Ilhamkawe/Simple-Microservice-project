package riviews

import "time"

type Riviews struct {
	ID        int
	UserID    int
	CourseID  int
	Rating    int
	Note      string
	CreatedAt time.Time
	UpdatedAt time.Time
}