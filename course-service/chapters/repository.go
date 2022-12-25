package chapters

import "gorm.io/gorm"

type Repository interface {
	Create(chapter Chapters) (Chapters, error)
	FindByCourseID(id int) ([]Chapters, error)
	FindByID(id int) (Chapters, error)
	Index() ([]Chapters, error)
	Destroy(id int) (bool, error)
	Update(chapter Chapters) (Chapters, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(chapter Chapters) (Chapters, error){

	err := r.db.Create(&chapter).Error

	if err != nil {
		return chapter, err 
	}

	return chapter, nil

}

func (r *repository) Update(chapter Chapters) (Chapters, error) {
	err := r.db.Save(&chapter).Error 

	if err != nil {
		return chapter, err 
	}

	return chapter, nil
}

func (r *repository) Index() ([]Chapters, error) {
	var chapter []Chapters
	err := r.db.Find(&chapter).Error

	if err != nil {
		return chapter, err 
	}

	return chapter, nil

}

func (r *repository) Show(id int) (Chapters, error) {

	var chapter Chapters

	err := r.db.Find(&chapter).Error

	if err != nil {
		return chapter, err 
	}

	return chapter, nil

}

func (r *repository) FindByCourseID(id int) ([]Chapters, error){
	var chapter []Chapters

	err := r.db.Where("course_id = ?", id).Find(&chapter).Error

	if err != nil {
		return chapter, err
	}

	return chapter, nil

}
func (r *repository) FindByID(id int) (Chapters, error){
	var chapter Chapters

	err := r.db.Where("id = ?", id).Find(&chapter).Error

	if err != nil {
		return chapter, err
	}

	return chapter, nil

}

func (r *repository) Destroy(id int) (bool, error) {

	err := r.db.Delete(&Chapters{}, id).Error

	if err != nil {
		return false, err
	}

	return true, nil

}

