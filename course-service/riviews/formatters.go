package riviews

import "time"

type RiviewFormatter struct {
	ID        int `json:"id"`
	CourseID  int `json:"course_id"`
	Rating    int `json:"rating"`
	Note      string `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatRiview(riviewParam Riviews) RiviewFormatter {
	riview := RiviewFormatter{}
	riview.ID = riviewParam.ID
	riview.CourseID = riviewParam.CourseID
	riview.Rating = riviewParam.Rating
	riview.Note = riviewParam.Note

	return riview
}

func FormatRiviews(riviewParam []Riviews) []RiviewFormatter{
	riviewsFormatter := []RiviewFormatter{}
	
	for _, riview := range riviewParam {
		riviewFormatter := FormatRiview(riview)
		riviewsFormatter = append(riviewsFormatter, riviewFormatter)
	}

	return riviewsFormatter

}