package course

import "gorm.io/gorm"

type Repository interface {
	Create(mentor Mentors) (Mentors, error)
	FindMentorByEmail(email string) (Mentors, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
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
