package service

import (
	userRepository "github.com/atakanyeniceli/payroll/models/user/repository"
)

type Service struct {
	Repo *userRepository.Repository
}

func NewService(repo *userRepository.Repository) *Service {
	return &Service{Repo: repo}
}
