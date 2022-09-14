package courses

import (
	"time"
)

type CoursesFormatter struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Certificate bool      `json:"certificate"`
	Thumbnail   string    `json:"thumbnail"`
	Price       int       `json:"price"`
	Status      string    `json:"status"`
	Type        string    `json:"type"`
	Level       string    `json:"level"`
	Description string    `json:"description"`
	MentorID    int       `json:"mentor_id"`
	Mentor 		CourseMentorFormatter `json:"mentor"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CoursesInputFormatter struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Certificate bool      `json:"certificate"`
	Thumbnail   string    `json:"thumbnail"`
	Price       int       `json:"price"`
	Status      string    `json:"status"`
	Type        string    `json:"type"`
	Level       string    `json:"level"`
	Description string    `json:"description"`
	MentorID    int       `json:"mentor_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CourseMentorFormatter struct { 
	ID int `json:"id"`
	Name string `json:"name"` 
	Profile string `json:"profile"`
}
func FormatCourse(course Courses) CoursesInputFormatter {
	formatter := CoursesInputFormatter{}
	formatter.ID = course.ID
	formatter.Name = course.Name
	formatter.Certificate = course.Certificate
	formatter.Thumbnail = course.Thumbnail 
	formatter.Price = course.Price 
	formatter.Status = course.Status 
	formatter.Type = course.Type 
	formatter.Level = course.Level 
	formatter.Description = course.Description 
	formatter.MentorID = course.MentorID

	return formatter
}

func ShowFormatCourses(course Courses) CoursesFormatter {
	formatter := CoursesFormatter{}
	formatter.ID = course.ID
	formatter.Name = course.Name
	formatter.Certificate = course.Certificate
	formatter.Thumbnail = course.Thumbnail 
	formatter.Price = course.Price 
	formatter.Status = course.Status 
	formatter.Type = course.Type 
	formatter.Level = course.Level 
	formatter.Description = course.Description 
	formatter.MentorID = course.MentorID

	mentor := course.Mentor

	mentorCourse := CourseMentorFormatter{}
	mentorCourse.ID = mentor.ID
	mentorCourse.Name = mentor.Name
	mentorCourse.Profile = mentor.Profile
	
	formatter.Mentor = mentorCourse

	return formatter
}

func FormatCourses(course []Courses) []CoursesFormatter {
	coursesFormatter := []CoursesFormatter{}
	for _, course := range course {
		courseFormatter := ShowFormatCourses(course)
		coursesFormatter = append(coursesFormatter, courseFormatter)
	}
	return coursesFormatter
}


// name
// certificate
// thumbnail
// price
// status
// type
// level
// description
// mentor_id