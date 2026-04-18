package repository

import (
	"context"

	"github.com/atakanyeniceli/payroll/models/overtime"
)

func (r *Repository) Update(ctx context.Context, o *overtime.Overtime) error {
	// UPDATE sorgusundan 'amount' kaldırıldı, parametre sıraları güncellendi
	query := `
		UPDATE overtimes 
		SET work_date = $1, duration_hours = $2, multiplier = $3, description = $4 
		WHERE id = $5 AND user_id = $6`

	_, err := r.DB.ExecContext(ctx, query, o.WorkDate, o.DurationHours, o.Multiplier, o.Description, o.ID, o.UserID)
	return err
}
