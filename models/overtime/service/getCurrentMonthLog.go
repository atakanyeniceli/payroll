package service

import (
	"context"
	"time"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/models/overtime"
)

func (s *Service) GetCurrentMonthLog(ctx context.Context, userID int) ([]*overtime.Overtime, error) {
	now := time.Now()

	// Mevcut ayın ilk günü (Yıl, Ay, 1, 00:00:00)
	startDate := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)

	// Bir sonraki ayın ilk günü (Sorguda bitiş sınırı olarak kullanılacak)
	endDate := startDate.AddDate(0, 1, 0)

	// Repository'den belirtilen tarih aralığındaki kayıtları getir
	overtimes, err := s.Repo.ListByDateRange(ctx, userID, startDate, endDate)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}

	return overtimes, nil
}
