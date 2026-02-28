package service

import (
	"github.com/atakanyeniceli/payroll/models/extra/repository"
)

type Service struct {
	Repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Repo: repo}
}
