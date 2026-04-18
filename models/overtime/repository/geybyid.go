package repository

import (
	"context"

	"github.com/atakanyeniceli/payroll/models/overtime"
)

func (r *Repository) GetByID(ctx context.Context, id int, userID int) (*overtime.Overtime, error) {
	// SELECT sorgusundan 'amount' kaldırıldı
	query := `
		SELECT id, user_id, work_date, start_time, end_time, duration_hours, multiplier, description 
		FROM overtimes 
		WHERE id = $1 AND user_id = $2`

	row := r.DB.QueryRowContext(ctx, query, id, userID)

	var o overtime.Overtime
	// Scan metodundan 'Amount' kaldırıldı
	err := row.Scan(&o.ID, &o.UserID, &o.WorkDate, &o.StartTime, &o.EndTime, &o.DurationHours, &o.Multiplier, &o.Description)
	if err != nil {
		return nil, err
	}
	if len(o.StartTime) > 16 && o.StartTime[10] == 'T' {
		o.StartTime = o.StartTime[11:16] // Sadece "14:51" kısmını alır
	}

	if len(o.EndTime) > 16 && o.EndTime[10] == 'T' {
		o.EndTime = o.EndTime[11:16] // Sadece "18:00" kısmını alır
	}

	return &o, nil
}
