package service

import (
	"context"

	apperror "github.com/atakanyeniceli/payroll/models/appError"
)

func (s *Service) Delete(ctx context.Context, id int, userID int) error {
	if err := s.Repo.Delete(ctx, id, userID); err != nil {
		return apperror.NewServerError(err)
	}
	return nil
}
