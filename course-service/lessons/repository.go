package lessons

import "gorm.io/gorm"

type Repository interface{
	Create(lesson Lessons) (Lessons, error)
	FindByID(id int) (Lessons, error)
	FindByChapterID(id int) ([]Lessons, error)
	GetAll() ([]Lessons, error)
	Destroy(id int) (bool, error)
	Update(lesson Lessons) (Lessons, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(lesson Lessons) (Lessons, error){

	err := r.db.Create(&lesson).Error

	if err != nil {
		return lesson, err
	}

	return lesson, nil

}

func (r *repository) Update(lesson Lessons) (Lessons, error){
	err := r.db.Save(&lesson).Error 

	if err != nil {
		return lesson, err
	}

	return lesson, nil
}

func (r *repository) FindByChapterID(id int) ([]Lessons, error){
	var lessons []Lessons
	err := r.db.Where("chapter_id = ?", id).Find(&lessons).Error
	
	if err != nil {
		return lessons, err
	}

	return lessons, nil 
}

func (r *repository) FindByID(id int) (Lessons, error) {
	var lesson Lessons
	err := r.db.Where("id = ?", id).Find(&lesson).Error

	if err != nil {
		return lesson, err
	}

	return lesson, nil 
}

func (r *repository) GetAll() ([]Lessons, error) {
	var lessons []Lessons
	err := r.db.Find(&lessons).Error 

	if err != nil {
		return lessons, err
	}

	return lessons, nil 
}

func (r *repository) Destroy(id int) (bool, error) {
	err := r.db.Delete(&Lessons{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}