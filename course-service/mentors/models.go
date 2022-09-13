package mentors

import "time"

type Mentors struct {
	ID         int
	Name       string
	Profile    string
	Email      string
	Profession string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}


// type MyCourses struct {
// 	ID        int
// 	CourseID  int
// 	Course    Courses
// 	UserID    int
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

// type Chapters struct {
// 	ID        int
// 	Name      string
// 	CourseID  int
// 	Course    Courses
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

// type Lessons struct {
// 	ID        int
// 	Name      string
// 	Video     string
// 	ChapterID int
// 	Chapter   Chapters
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

// type ImageCourses struct {
// 	ID        int
// 	CourseID  int
// 	Course    Courses
// 	Image     string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

// type Riviews struct {
// 	ID        int
// 	UserID    int
// 	CourseID  int
// 	Course    Courses
// 	Rating    int
// 	Note      string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }
