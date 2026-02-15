package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (repo *Repository) GetByDate(ctx context.Context, userid int, date time.Time) (float64, error) {
	query := `
		SELECT amount 
		FROM hourly_rates 
		WHERE user_id = $1 AND effective_date <= $2 
		ORDER BY effective_date DESC 
		LIMIT 1
	`

	var amount float64
	err := repo.DB.QueryRowContext(ctx, query, userid, date).Scan(&amount)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("saatlik ücret getirilirken veritabanı hatası: %w", err)
	}

	return amount, nil
}
