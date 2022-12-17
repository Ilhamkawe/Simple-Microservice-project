package lessons

import "time"

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