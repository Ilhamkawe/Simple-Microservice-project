package orders

import (
	"time"

	"gorm.io/datatypes"
)
type Orders struct {
	ID        int
	Status    string
	UserID   int
	CourseID int
	SnapUrl  string
	Metadata  datatypes.JSON
	CreatedAt time.Time
	UpdatedAt time.Time
}