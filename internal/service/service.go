package service

import (
	"context"
	"projectWithPostgre/internal/repository"
)

type Servicer interface {
	GetUser(ctx context.Context, id string) (repository.User, error)
	CreateUser(ctx context.Context, user repository.User) error
	UpdateUser(ctx context.Context, user repository.User) error
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, c repository.Conditions) ([]repository.User, int, error)
}
type Service struct {
	repository.Repository
}

func New(repository2 *repository.Repository) *Service {
	return &Service{
		*repository2,
	}
}

func (s *Service) GetUser(ctx context.Context, id string) (repository.User, error) {
	return s.Repository.GetByID(ctx, id)
}
func (s *Service) CreateUser(ctx context.Context, user repository.User) error {
	return s.Repository.Create(ctx, user)
}
func (s *Service) UpdateUser(ctx context.Context, user repository.User) error {
	return s.Repository.Update(ctx, user)
}
func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.Repository.Delete(ctx, id)
}
func (s *Service) ListUsers(ctx context.Context, c repository.Conditions) ([]repository.User, int, error) {
	return s.Repository.List(ctx, c)
}
