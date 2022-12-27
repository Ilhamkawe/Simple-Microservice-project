package courses

import "gorm.io/gorm"

type Repository interface {
	Create(course Courses)(Courses, error)
	Update(course Courses) (Courses, error)
	FindCourseByID(id int) (Courses, error)
	GetAll()([]Courses, error)
	Destroy(id int)(bool, error)
	FilterCourse(name string, type_ string)([]Courses, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}


func (r *repository) Destroy(id int)(bool, error){
	err := r.db.Delete(&Courses{}, id).Error
	if err != nil { 
		return false, err
	}
	return true, nil
}


func (r *repository) Create(course Courses)(Courses, error) {
	err := r.db.Create(&course).Error

	if err != nil {
		return course, err
	}

	return course, nil
}

func (r *repository) FindCourseByID(id int) (Courses, error){
	var course Courses
	err := r.db.Where("id = ?", id).Preload("Mentor").Preload("Chapters.Lessons").Find(&course).Error

	if err != nil {
		return course, err
	}

	return course, nil
}

func (r *repository) Update(course Courses) (Courses, error){
	err := r.db.Save(&course).Error

	if err != nil {
		return course, err 
	}

	return course, nil
}

func (r *repository) GetAll()([]Courses, error){
	var courses []Courses 
	err := r.db.Preload("Mentor").Find(&courses).Error
	if err != nil { 
		return courses, err
	}

	return courses, nil 
}

func (r *repository) FilterCourse(name string, type_ string)([]Courses, error){
	var courses []Courses 
	var err error
	if name != "" {
		if type_ != ""{
			err = r.db.Preload("Mentor").Where("name LIKE ?", "%"+name+"%").Where("type = ?" , type_).Find(&courses).Error
		}else {
			err = r.db.Preload("Mentor").Where("name LIKE  ?", "%"+name+"%").Find(&courses).Error
		}
	}else{
		err = r.db.Preload("Mentor").Where("type = ?" , type_).Find(&courses).Error
	}
	if err != nil { 
		return courses, err
	}
	return courses, nil 
}