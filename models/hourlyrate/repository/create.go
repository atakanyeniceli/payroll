package repository

import (
	"context"
	"fmt"

	"github.com/atakanyeniceli/payroll/models/hourlyrate"
)

func (r *Repository) Create(ctx context.Context, rate *hourlyrate.HourlyRate) error {
	query := `INSERT INTO hourly_rates (user_id, amount, effective_date) VALUES ($1, $2, $3) RETURNING id`

	err := r.DB.QueryRowContext(ctx, query, rate.UserID, rate.Amount, rate.EffectiveDate).Scan(&rate.ID)
	if err != nil {
		return fmt.Errorf("saat ücreti kaydedilirken veritabanı hatası: %w", err)
	}

	return nil
}
