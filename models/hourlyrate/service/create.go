package service

import (
	"context"
	"strconv"
	"strings"
	"time"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/models/hourlyrate"
)

type CreateHourlyRateDTO struct {
	UserID           int
	AmountStr        string
	EffectiveDateStr string
}

func (s *Service) Create(ctx context.Context, dto CreateHourlyRateDTO) (*hourlyrate.HourlyRate, error) {
	// 0. Timeout ve Temizlik
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	dto.AmountStr = strings.TrimSpace(dto.AmountStr)
	dto.EffectiveDateStr = strings.TrimSpace(dto.EffectiveDateStr)

	// 1. Validasyon
	if dto.AmountStr == "" || dto.EffectiveDateStr == "" {
		return nil, apperror.NewClientError("Lütfen tüm alanları doldurunuz.", 400)
	}

	// 2. Veri Dönüşümü
	amount, err := strconv.ParseFloat(dto.AmountStr, 64)
	if err != nil {
		return nil, apperror.NewClientError("Geçersiz ücret formatı.", 400)
	}

	effectiveDate, err := time.Parse("2006-01-02", dto.EffectiveDateStr)
	if err != nil {
		return nil, apperror.NewClientError("Geçersiz tarih formatı.", 400)
	}

	// 3. Model Oluşturma ve Kayıt
	rate := &hourlyrate.HourlyRate{
		UserID:        dto.UserID,
		Amount:        amount,
		EffectiveDate: effectiveDate,
	}

	if err := s.Repo.Create(ctx, rate); err != nil {
		return nil, apperror.NewServerError(err)
	}

	return rate, nil
}
