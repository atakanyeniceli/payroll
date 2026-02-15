package service

import (
	"context"
	"strconv"
	"time"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/models/overtime"
)

// DTO: Web formundan veya JSON'dan gelen ham veriyi burada topluyoruz
type CreateOvertimeDTO struct {
	UserID      int
	DateStr     string // "2023-10-25"
	StartTime   string // "18:00"
	EndTime     string // "20:00"
	Multiplier  string // "1.5"
	Description string
}

func (s *Service) CreateOvertime(ctx context.Context, dto CreateOvertimeDTO) (*overtime.Overtime, error) {
	// 0. Timeout Tanımla
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// 1. Validasyonlar
	if dto.DateStr == "" || dto.StartTime == "" || dto.EndTime == "" {
		return nil, apperror.NewClientError("Tarih ve saat alanları zorunludur.", 400)
	}

	// 2. Tarih Parse Et
	workDate, err := time.Parse("2006-01-02", dto.DateStr)
	if err != nil {
		return nil, apperror.NewClientError("Geçersiz tarih formatı.", 400)
	}

	// 3. Saat Kontrolü (Mantıksal)
	t1, err1 := time.Parse("15:04", dto.StartTime)
	t2, err2 := time.Parse("15:04", dto.EndTime)
	if err1 != nil || err2 != nil {
		return nil, apperror.NewClientError("Saat formatı geçersiz (HH:MM).", 400)
	}
	if !t2.After(t1) {
		return nil, apperror.NewClientError("Bitiş saati başlangıçtan sonra olmalıdır.", 400)
	}

	// 4. Çarpan (Multiplier) Parse Et
	multiplier, err := strconv.ParseFloat(dto.Multiplier, 64)
	if err != nil {
		multiplier = 1.50 // Varsayılan değer (Hata fırlatmıyoruz, varsayılan atıyoruz)
	}

	// 5. Model Oluştur
	newOvertime := &overtime.Overtime{
		UserID:      dto.UserID,
		WorkDate:    workDate,
		StartTime:   dto.StartTime,
		EndTime:     dto.EndTime,
		Multiplier:  multiplier,
		Description: dto.Description,
	}

	// 6. Veritabanına Yaz (Repo)
	if err := s.Repo.Create(ctx, newOvertime); err != nil {
		return nil, apperror.NewServerError(err)
	}

	return newOvertime, nil
}
