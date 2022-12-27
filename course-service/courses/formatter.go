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

type CoursesDetailFormatter struct {
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
	Chapters 	[]ChaptersFormatter `json:"chapter"`
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

type ChaptersFormatter struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CourseID int `json:"course_id,omitempty"`
	Lessons []LessonFormatter `json:"lessons,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func FormatShowChapters(chapter Chapters) ChaptersFormatter {
	
	formatter := ChaptersFormatter{}
	formatter.ID = chapter.ID
	formatter.Name = chapter.Name
	formatter.CourseID = chapter.CourseID

	formatter.Lessons = FormatLessons(chapter.Lessons)

	return formatter
}

func FormatAllChapters(chapters []Chapters) []ChaptersFormatter {
	chaptersFormatter := []ChaptersFormatter{}
	for _, chapter := range chapters {
		chapterFormatter := FormatShowChapters(chapter)
		chaptersFormatter = append(chaptersFormatter, chapterFormatter)
	}

	return chaptersFormatter
}

type LessonFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Video     string    `json:"video"`
	ChapterID int    `json:"chapter_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatLesson(lesson Lessons) LessonFormatter {
	formatter := LessonFormatter{}
	formatter.ID = lesson.ID 
	formatter.Name = lesson.Name 
	formatter.Video = lesson.Video
	formatter.ChapterID = lesson.ChapterID
	formatter.CreatedAt = lesson.CreatedAt
	formatter.UpdatedAt = lesson.UpdatedAt

	return formatter
}

func FormatLessons(lesson []Lessons) []LessonFormatter {
	lessonsFormatter := []LessonFormatter{}

	for _, lesson := range lesson {
		lessonFormatter := FormatLesson(lesson)
		lessonsFormatter = append(lessonsFormatter, lessonFormatter)
	}
	return lessonsFormatter
}

func DetailFormatCourses(course Courses) CoursesDetailFormatter {
	formatter := CoursesDetailFormatter{}
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
	
	chapter := course.Chapters
	
	chapterCourse := FormatAllChapters(chapter)
	
	formatter.Chapters = chapterCourse
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