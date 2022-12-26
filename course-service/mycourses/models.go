package mycourse

import (
	"course-service/courses"
	"time"
)

type MyCourses struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	CourseID  int `json:"course_id"`
	Course courses.Courses `json:"course"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}