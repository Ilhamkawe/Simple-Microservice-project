package courses

import "gorm.io/gorm"

type Repository interface {
	Create(course Courses)(Courses, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(course Courses)(Courses, error) {
	err := r.db.Create(&course).Error

	if err != nil {
		return course, err
	}

	return course, nil
}