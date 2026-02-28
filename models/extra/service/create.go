package service

import (
	"context"
	"net/http"
	"time"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/models/extra"
)

// Create: Yeni bir ek ödeme oluşturur
func (s *Service) Create(ctx context.Context, userID int, title string, amount float64, date time.Time) error {
	// 1. Validasyonlar (Mevcut NewClientError kullanılarak)
	if amount <= 0 {
		return apperror.NewClientError("Tutar 0'dan büyük olmalıdır.", http.StatusBadRequest)
	}
	if title == "" {
		return apperror.NewClientError("Açıklama boş olamaz.", http.StatusBadRequest)
	}

	// 2. Model Hazırlığı
	newExtra := &extra.Extra{
		UserID: userID,
		Title:  title,
		Amount: amount,
		Date:   date,
	}

	// 3. Kayıt
	if err := s.Repo.Create(ctx, newExtra); err != nil {
		return apperror.NewServerError(err)
	}

	return nil
}
