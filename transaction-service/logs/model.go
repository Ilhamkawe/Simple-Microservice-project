package logs

import (
	"time"

	"gorm.io/datatypes"
)

type PaymentLogs struct {
	ID          int
	Status      string
	PaymentType string
	OrderID     int
	RawResponse datatypes.JSON
	CreatedAt time.Time
	UpdatedAt time.Time
}