package repository

import (
	"context"

	"github.com/atakanyeniceli/payroll/models/overtime"
)

func (r *Repository) Create(ctx context.Context, ot *overtime.Overtime) error {
	// duration_hours alanını insert etmiyoruz, DB hesaplıyor.
	query := `
		INSERT INTO overtimes (
			user_id, work_date, start_time, end_time, multiplier, description
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, duration_hours
	`

	// Context ile sorguyu çalıştırıyoruz
	err := r.DB.QueryRowContext(ctx, query,
		ot.UserID,
		ot.WorkDate,
		ot.StartTime,
		ot.EndTime,
		ot.Multiplier,
		ot.Description,
	).Scan(&ot.ID, &ot.DurationHours) // Hesaplanan süreyi geri al

	return err
}
