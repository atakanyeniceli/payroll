package summary

type SummaryStats struct {
	BaseSalary       float64 // Çıplak Maaş (Gün * 7.5 * Saat Ücreti)
	OvertimeEarnings float64 // Fazla Mesai Kazancı
	TotalEarnings    float64 // Hepsini Toplamı (Cebine girecek para)
}
