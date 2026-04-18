package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/atakanyeniceli/payroll/models/hourlyrate"
)

func (repo *Repository) GetByDate(ctx context.Context, userid int, date time.Time) (hourlyrate.HourlyRate, error) {
	query := `
		SELECT id,amount,effective_date
		FROM hourly_rates 
		WHERE user_id = $1 AND effective_date <= $2 
		ORDER BY effective_date DESC 
		LIMIT 1
	`

	var h hourlyrate.HourlyRate
	err := repo.DB.QueryRowContext(ctx, query, userid, date).Scan(&h.ID, &h.Amount, &h.EffectiveDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return h, nil
		}
		return h, fmt.Errorf("saatlik ücret getirilirken veritabanı hatası: %w", err)
	}

	return h, nil
}
