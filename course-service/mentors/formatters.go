package mentors

import "time"

type CourseFormatter struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	Profile   string    `json:"profile"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatCourse(mentor Mentors) CourseFormatter {
	formatter := CourseFormatter{}
	formatter.ID = mentor.ID
	formatter.Name = mentor.Name
	formatter.Email = mentor.Email
	formatter.Profile = mentor.Profile
	return formatter
}

func FormatMentors(mentors []Mentors) []CourseFormatter {
	mentorsFormatter := []CourseFormatter{}
	for _, mentor := range mentors {
		courseFormatter := FormatCourse(mentor)
		mentorsFormatter = append(mentorsFormatter, courseFormatter)
	} 

	return mentorsFormatter
}
