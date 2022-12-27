package chapters

import (
	"course-service/courses"
	"time"
)

type Chapters struct {
	ID       int 
	Name     string 
	CourseID int 
	Course   courses.Courses 
	CreatedAt time.Time 
	UpdatedAt time.Time  
}
