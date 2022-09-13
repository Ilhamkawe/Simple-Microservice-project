package mentors

import "gorm.io/gorm"

type Repository interface {
	Create(mentor Mentors) (Mentors, error)
	FindMentorByEmail(email string) (Mentors, error)
	Update(mentor Mentors) (Mentors, error)
	FindMentorByID(id int) (Mentors, error)
	GetAllMentor()([]Mentors, error)
	Destroy(id int)(bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Destroy(id int)(bool, error) {
	err := r.db.Delete(&Mentors{}, id).Error

	if err != nil { 
		return false, err
	}

	return true, nil
}

// cari semua data mentor 
func (r *repository) GetAllMentor()([]Mentors, error) { 
	var mentor []Mentors
	err := r.db.Find(&mentor).Error

	if err != nil { 
		return mentor, err
	}

	return mentor, nil 
}

// input data mentors
func (r *repository) Create(mentor Mentors) (Mentors, error) {
	err := r.db.Create(&mentor).Error
	if err != nil {
		return mentor, err
	}

	return mentor, err
}

// cari data mentor berdasarkan email
func (r *repository) FindMentorByEmail(email string) (Mentors, error) {
	var mentor Mentors
	err := r.db.Where("email = ?", email).Find(&mentor).Error

	if err != nil {
		return mentor, err
	}

	return mentor, nil
}

// cari data mentor berdasarkan id
func (r *repository) FindMentorByID(id int) (Mentors, error) {
	var mentor Mentors 
	err := r.db.Where("id = ?", id).Find(&mentor).Error

	if err != nil {
		return mentor, err 
	}

	return mentor, nil
}

// update mentors 
func (r *repository) Update(mentor Mentors) (Mentors, error) {
		err := r.db.Save(&mentor).Error

		if err != nil {
			return mentor, err
		}

		return mentor, nil
}
