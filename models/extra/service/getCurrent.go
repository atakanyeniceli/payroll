package service

import (
	"context"
	"time"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/models/extra"
)

func (s *Service) GetCurrent(ctx context.Context, userID int) ([]*extra.Extra, error) {
	now := time.Now()

	// Mevcut ayın ilk günü
	startDate := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)

	// Bir sonraki ayın ilk günü (Bitiş sınırı)
	endDate := startDate.AddDate(0, 1, 0)

	extras, err := s.Repo.ListByDateRange(ctx, userID, startDate, endDate)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	return extras, nil
}
