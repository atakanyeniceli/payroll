package service

import (
	hourlyRateRepository "github.com/atakanyeniceli/payroll/models/hourlyrate/repository"
)

type Service struct {
	Repo *hourlyRateRepository.Repository
}

func NewService(repo *hourlyRateRepository.Repository) *Service {
	return &Service{Repo: repo}
}
