package tasks

import (
	"context"
	model "token-based-auth/internal/domain"
)

type Repository interface {
	Get(ctx context.Context, tID string) *model.Task
	Create(ctx context.Context, i model.Info) *model.Task
}

type Service struct {
	repo Repository
}

func New(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Get(ctx context.Context, tID string) *model.Task {
	return s.repo.Get(ctx, tID)
}

func (s *Service) Create(ctx context.Context, i model.Info) *model.Task {
	return s.repo.Create(ctx, i)
}
