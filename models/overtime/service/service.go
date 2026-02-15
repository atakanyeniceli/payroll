package service

import (
	hourlyRateService "github.com/atakanyeniceli/payroll/models/hourlyrate/service"
	overtimeRepository "github.com/atakanyeniceli/payroll/models/overtime/repository"
)

type Service struct {
	Repo    *overtimeRepository.Repository
	RateSrv *hourlyRateService.Service
}

func NewService(repo *overtimeRepository.Repository, rateSrv *hourlyRateService.Service) *Service {
	return &Service{
		Repo:    repo,
		RateSrv: rateSrv,
	}
}
