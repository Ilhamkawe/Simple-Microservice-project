package riviews

import "gorm.io/gorm"

type Repository interface {
	Create(riview Riviews) (Riviews, error)
	FindByUserAndCourseID(course_id int, user_id int) (Riviews, error)
	Update(riview Riviews) (Riviews, error)
	Destroy(id int) (bool, error)
	GetByID(id int) (Riviews, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(riview Riviews) (Riviews, error) {
	err := r.db.Create(&riview).Error 

	if err != nil {
		return riview, err
	}
	
	return riview, nil
}

func (r *repository) FindByUserAndCourseID(course_id int, user_id int) (Riviews, error) {
	var riview Riviews

	err := r.db.Where("user_id = ? AND course_id = ?", user_id, course_id).Find(&riview).Error 

	if err != nil {
		return riview, err
	}

	return riview, nil
}

func (r *repository) Update(riview Riviews) (Riviews, error){
	err := r.db.Save(&riview).Error

	if err != nil {
		return riview, err 
	}

	return riview, nil
}

func (r *repository) Destroy(id int) (bool, error) {
	err := r.db.Delete(&Riviews{},id).Error

	if err != nil {
		return false, err 
	}

	return true, nil
}

func (r *repository) GetByID(id int) (Riviews, error) {
	var riview Riviews
	err := r.db.Where("id = ?", id).Find(&riview).Error

	if err != nil {
		return riview, err
	}

	return riview, nil
}