package extra

import "time"

type Extra struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}
