package mycourse

import "gorm.io/gorm"

type Repository interface{
	FindByUserID(id int) ([]MyCourses, error)
	GetAll() ([]MyCourses, error)
	Create(mycourse MyCourses) (MyCourses, error)
	FindByUserAndCourseID(userID int, courseID int) (MyCourses, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(mycourse MyCourses) (MyCourses, error) {
	err := r.db.Create(&mycourse).Error

	if err != nil {
		return mycourse, err 
	}

	return mycourse, nil
}

func (r *repository) GetAll() ([]MyCourses, error) {
	var mycourses []MyCourses
	err := r.db.Find(&mycourses).Error

	if err != nil {
		return mycourses, err
	}

	return mycourses, nil
}

func (r *repository) FindByUserID(id int) ([]MyCourses, error) {
	var mycourses []MyCourses

	err := r.db.Where("user_id = ?", id).Find(&mycourses).Error

	if err != nil {
		return mycourses, err
	}

	return mycourses, nil
}

func (r *repository) FindByUserAndCourseID(userID int, courseID int) (MyCourses, error) {
	var  myCourse MyCourses
	err := r.db.Where("course_id = ? AND user_id = ?", courseID, userID).Find(&myCourse).Error

	if err != nil {
		return myCourse, err 
	}

	return myCourse, nil

}