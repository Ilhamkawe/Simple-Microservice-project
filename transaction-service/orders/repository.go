package orders

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Repository interface {
	Create(order Orders) (Orders, error)
	Update(order Orders) (Orders, error)
	UpdateMetaData(id int, Metadata MetadataInput) (Orders, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(order Orders) (Orders, error){
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository) GetByID(ID int) (Orders, error) {
	
	var order Orders

	err := r.db.Where("id = ?", ID).Find(&order).Error 

	if err != nil {
		return order, err 
	}

	return order, nil

}

func (r *repository) Update(order Orders) (Orders, error) {
	err := r.db.Save(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository) UpdateMetaData(id int, Metadata MetadataInput) (Orders, error) {
	order := Orders{}
	err := r.db.Model(&order).Where("id = ?", id).UpdateColumn("metadata", datatypes.JSONSet("medatadata").Set("course_id", Metadata.CourseID).Set("course_price", Metadata.CoursePrice).Set("course_name", Metadata.CourseName).Set("course_thumbnail", Metadata.CourseThumbnail).Set("course_level", Metadata.CourseLevel)).Error
	if err != nil {
		return order, err
	}

	return order, nil
}




