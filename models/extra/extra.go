package extra

import "time"

type Extra struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Title  string    `json:"title"`
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
}
