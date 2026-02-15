package service

import (
	"context"
	"time"
)

func (s *Service) GetByIDAndDate(ctx context.Context, userID int, date time.Time) (float64, error) {
	return s.Repo.GetByDate(ctx, userID, date)
}
