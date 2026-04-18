package service

import (
	"context"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
	"github.com/atakanyeniceli/payroll/models/overtime"
)

func (s *Service) GetByID(ctx context.Context, id, userID int) (*overtime.Overtime, error) {
	o, err := s.Repo.GetByID(ctx, id, userID)
	if err != nil {
		return nil, apperror.NewServerError(err)
	}
	return o, nil
}
