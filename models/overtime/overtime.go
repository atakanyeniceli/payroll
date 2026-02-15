package overtime

import (
	"time"
)

type Overtime struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	WorkDate      time.Time `json:"work_date"`      // DB: work_date (DATE)
	StartTime     string    `json:"start_time"`     // DB: start_time (TIME) - "18:30"
	EndTime       string    `json:"end_time"`       // DB: end_time (TIME) - "20:00"
	Multiplier    float64   `json:"multiplier"`     // DB: multiplier (NUMERIC) - 1.50
	DurationHours float64   `json:"duration_hours"` // DB: Generated (Read-Only)
	Amount        float64   `json:"amount"`         // DB: Generated (Read-Only)
	Description   string    `json:"description"`
}

func NewOvertime() *Overtime {
	return &Overtime{}
}
