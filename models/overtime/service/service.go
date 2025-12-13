package service

import (
	overtimeRepository "github.com/atakanyeniceli/payroll/models/overtime/repository"
)

type Service struct {
	Repo *overtimeRepository.Repository
}

func NewService(repo *overtimeRepository.Repository) *Service {
	return &Service{Repo: repo}
}
