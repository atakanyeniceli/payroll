package repository

import (
	"context"

	"github.com/atakanyeniceli/payroll/models/extra"
)

func (r *Repository) Create(ctx context.Context, e *extra.Extra) error {
	query := `
		INSERT INTO extras (user_id, title, amount, payment_date) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, created_at`

	err := r.DB.QueryRowContext(ctx, query, e.UserID, e.Title, e.Amount, e.Date).
		Scan(&e.ID)

	return err
}
