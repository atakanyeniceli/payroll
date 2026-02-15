package service

import (
	hourlyRateService "github.com/atakanyeniceli/payroll/models/hourlyrate/service"
	overtimeService "github.com/atakanyeniceli/payroll/models/overtime/service"
	// extraService "github.com/atakanyeniceli/payroll/models/extra/service"
)

type Service struct {
	// Repository YOK! Onun yerine diğer servisler var.
	OvertimeSrv *overtimeService.Service
	RateSrv     *hourlyRateService.Service // Maaş hesabı için Rate servisine ihtiyacımız var

}

// Constructor: Bağımlılıkları (Diğer servisleri) alıyor
func NewService(overtimeSrv *overtimeService.Service, rateSrv *hourlyRateService.Service) *Service {
	return &Service{
		OvertimeSrv: overtimeSrv,
		RateSrv:     rateSrv,
	}
}
