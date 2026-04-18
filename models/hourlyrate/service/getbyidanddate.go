package service

import (
	"context"
	"time"

	"github.com/atakanyeniceli/payroll/models/hourlyrate"
)

func (s *Service) GetByIDAndDate(ctx context.Context, userID int, date time.Time) (hourlyrate.HourlyRate, error) {
	return s.Repo.GetByDate(ctx, userID, date)
}
