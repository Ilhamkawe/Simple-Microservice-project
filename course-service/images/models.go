package images

import "time"

type ImageCourses struct {
	ID        int
	CourseID  int
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}