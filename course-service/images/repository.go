package images

import "gorm.io/gorm"

type Repository interface {
	Create(img ImageCourses) (ImageCourses, error)
	Destroy(id int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(img ImageCourses) (ImageCourses, error) {

	err := r.db.Create(&img).Error

	if err != nil {
		return img, err 
	}

	return img, err

}

func (r*repository) Destroy(id int) (bool, error) {
	err := r.db.Delete(&ImageCourses{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}