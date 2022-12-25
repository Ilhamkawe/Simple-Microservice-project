package chapters

import "time"

type ChaptersFormatter struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CourseID int `json:"course_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func FormatShowChapters(chapter Chapters) ChaptersFormatter {
	
	formatter := ChaptersFormatter{}
	formatter.ID = chapter.ID
	formatter.Name = chapter.Name
	formatter.CourseID = chapter.CourseID

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