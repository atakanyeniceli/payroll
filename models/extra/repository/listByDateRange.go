package repository

import (
	"context"
	"time"

	"github.com/atakanyeniceli/payroll/models/extra"
)

func (r *Repository) ListByDateRange(ctx context.Context, userID int, startDate, endDate time.Time) ([]*extra.Extra, error) {
	query := `
		SELECT id, title, amount, payment_date 
		FROM extras 
		WHERE user_id = $1 AND payment_date  >= $2 AND payment_date  < $3 
		ORDER BY payment_date  ASC`

	rows, err := r.DB.QueryContext(ctx, query, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var extras []*extra.Extra
	for rows.Next() {
		e := &extra.Extra{UserID: userID}
		if err := rows.Scan(&e.ID, &e.Title, &e.Amount, &e.Date); err != nil {
			return nil, err
		}
		extras = append(extras, e)
	}

	return extras, rows.Err()
}
