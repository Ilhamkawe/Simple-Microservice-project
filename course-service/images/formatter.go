package images

import "time"

type ImageCoursesFormatter struct {
	ID        int       `json:"id,omitempty"`
	CourseID int       `json:"chapter_id"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatImageCourse(img ImageCourses) ImageCoursesFormatter {
	formatter := ImageCoursesFormatter{}
	formatter.ID = img.ID
	formatter.Image = img.Image
	formatter.CourseID = img.CourseID
	formatter.CreatedAt = img.CreatedAt
	formatter.UpdatedAt = img.UpdatedAt

	return formatter
}

func FormatImageCourses(img []ImageCourses) []ImageCoursesFormatter{
	imagesFormatter := []ImageCoursesFormatter{}
	for _, img := range img {
		imgFormatter := FormatImageCourse(img)
		imagesFormatter = append(imagesFormatter, imgFormatter)
	}

	return imagesFormatter
}