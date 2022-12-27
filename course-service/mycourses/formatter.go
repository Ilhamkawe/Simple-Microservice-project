package mycourse

import "time"

type MyCourseFormatter struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	CourseID  int       `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatMyCourse(mycourseParam MyCourses) MyCourseFormatter {
	myCourse := MyCourseFormatter{}
	myCourse.ID = mycourseParam.ID
	myCourse.UserID = mycourseParam.UserID
	myCourse.CourseID = mycourseParam.UserID
	myCourse.CreatedAt = mycourseParam.CreatedAt
	myCourse.UpdatedAt = mycourseParam.UpdatedAt
	return myCourse
}

func FormatMyCourses(myCourseParam []MyCourses) []MyCourseFormatter {
	myCoursesFormatter := []MyCourseFormatter{}
	for _, mycourse := range myCourseParam {
		myCourseFormatter := FormatMyCourse(mycourse)
		myCoursesFormatter = append(myCoursesFormatter, myCourseFormatter) 
	}

	return myCoursesFormatter
}