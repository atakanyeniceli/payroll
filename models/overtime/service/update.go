package service

import (
	"context"
	"net/http"
	"time"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/models/overtime"
)

func (s *Service) Update(ctx context.Context, id, userID int, date time.Time, startTimeStr, endTimeStr string, multiplier float64, description string) error {

	// 1. Saatleri kontrol et ve Süreyi hesapla
	startTime, err := time.Parse("15:04", startTimeStr)
	if err != nil {
		return apperror.NewClientError("Geçersiz başlangıç saati", http.StatusBadRequest)
	}
	endTime, err := time.Parse("15:04", endTimeStr)
	if err != nil {
		return apperror.NewClientError("Geçersiz bitiş saati", http.StatusBadRequest)
	}

	duration := endTime.Sub(startTime).Hours()

	// Gece yarısını geçme (Overnight) kontrolü
	if duration < 0 {
		duration += 24.0
	}
	// Mantıksız girişleri engelle
	if duration > 16 {
		return apperror.NewClientError("Bir mesai tek seferde 16 saatten uzun olamaz.", http.StatusBadRequest)
	}
	if duration == 0 {
		return apperror.NewClientError("Başlangıç ve bitiş saati aynı olamaz.", http.StatusBadRequest)
	}

	// 2. Modeli Güncelle (Amount olmadan)
	updatedOvertime := &overtime.Overtime{
		ID:            id,
		UserID:        userID,
		WorkDate:      date,
		DurationHours: duration,
		Multiplier:    multiplier,
		Description:   description,
	}

	// 3. Veritabanına kaydet
	if err := s.Repo.Update(ctx, updatedOvertime); err != nil {
		return apperror.NewServerError(err)
	}

	return nil
}
