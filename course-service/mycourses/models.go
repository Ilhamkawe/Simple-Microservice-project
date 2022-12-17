package mycourse

import "time"

type MyCourses struct {
	ID        int `json:"id"`
	CourseID  int `json:"course_id"`
	UserID    int `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}