package orders

import (
	"time"

	"gorm.io/datatypes"
)
type Orders struct {
	ID        int
	Status    string
	user_id   int
	course_id int
	snap_url  string
	metadata  datatypes.JSON
	CreatedAt time.Time
	UpdatedAt time.Time
}