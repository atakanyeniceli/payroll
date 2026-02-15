package service

import (
	"context"
	"time"

	"github.com/atakanyeniceli/payroll/models/summary"
)

// Helper: Ayın kaç gün çektiğini bulur (Şubat kuralı dahil)
func (s *Service) getSalaryDays(year int, month time.Month) int {
	if month == time.February {
		return 30 // Şubat kuralı: Daima 30 gün
	}
	// Diğer aylar için takvimdeki gerçek gün sayısı (30 veya 31)
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func (s *Service) GetCurrent(ctx context.Context, userID int) (*summary.SummaryStats, error) {
	stats := &summary.SummaryStats{}
	now := time.Now()

	// 1. GÜNCEL SAATLİK ÜCRETİ BUL
	// Maaş hesabı için o anki geçerli ücreti alıyoruz
	currentRate, err := s.RateSrv.GetByIDAndDate(ctx, userID, now)
	if err != nil {
		return nil, err
	}

	// 2. ÇIPLAK MAAŞ HESABI (Base Salary)
	// Formül: Saat Ücreti * 7.5 * Gün Sayısı (Şubat=30)
	daysInMonth := s.getSalaryDays(now.Year(), now.Month())
	stats.BaseSalary = currentRate * 7.5 * float64(daysInMonth)

	// 3. FAZLA MESAİ KAZANCI
	overtimeLogs, err := s.OvertimeSrv.GetCurrentMonthLog(ctx, userID)
	if err == nil {
		for _, log := range overtimeLogs {
			// Sadece parayı topluyoruz, saati summary'de göstermeyeceğiz
			stats.OvertimeEarnings += log.Amount
		}
	}

	// 4. TOPLAM HAKEDİŞ
	// İleride buraya + ExtraEarnings eklenecek
	stats.TotalEarnings = stats.BaseSalary + stats.OvertimeEarnings

	return stats, nil
}
