package hourlyrate

import "time"

type HourlyRate struct {
	ID            int
	UserID        int
	Amount        float64
	EffectiveDate time.Time
	CreatedAt     time.Time
}
