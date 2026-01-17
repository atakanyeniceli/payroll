package repository

import (
	"context"
	"time"

	"github.com/atakanyeniceli/payroll/models/overtime"
)

func (r *Repository) ListByDateRange(ctx context.Context, userID int, startDate, endDate time.Time) ([]*overtime.Overtime, error) {
	query := `
		SELECT id, user_id, work_date, start_time, end_time, multiplier, description, duration_hours
		FROM overtimes
		WHERE user_id = $1 AND work_date >= $2 AND work_date < $3
		ORDER BY work_date DESC
	`

	rows, err := r.DB.QueryContext(ctx, query, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var overtimes []*overtime.Overtime
	for rows.Next() {
		var ot overtime.Overtime
		if err := rows.Scan(
			&ot.ID,
			&ot.UserID,
			&ot.WorkDate,
			&ot.StartTime,
			&ot.EndTime,
			&ot.Multiplier,
			&ot.Description,
			&ot.DurationHours,
		); err != nil {
			return nil, err
		}
		overtimes = append(overtimes, &ot)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return overtimes, nil
}
