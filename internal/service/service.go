package service

import (
	"context"
	"projectWithPostgre/internal/repository"
)

type Servicer interface {
	CreateUser(ctx context.Context, user repository.User) (repository.User, error)
	GetUser(ctx context.Context, name string) (repository.User, error)
	UpdateUser(ctx context.Context, name string) (repository.User, error)
	DeleteUser(ctx context.Context, name string) error
	StoreOrder(ctx context.Context, order repository.Order) (repository.Order, error)
	GetOrder(ctx context.Context, id int) (repository.Order, error)
	DeleteOrder(ctx context.Context, id int) error
	CreatePet(ctx context.Context, pet repository.Pet) (repository.Pet, error)
	UpdatePet(ctx context.Context, pet repository.Pet) (repository.Pet, error)
	FindPetByStatus(ctx context.Context, status string) (repository.Pet, error)
	FindPetById(ctx context.Context, id int) (repository.Pet, error)
	DeletePet(ctx context.Context, id int) error
}
type Service struct {
	repository.Repository
}

func New(repository2 *repository.Repository) *Service {
	return &Service{
		*repository2,
	}
}
func (s *Service) CreateUser(ctx context.Context, user repository.User) (repository.User, error) {
	return s.Repository.CreateUser(ctx, user)
}
func (s *Service) GetUser(ctx context.Context, name string) (repository.User, error) {
	return s.Repository.GetUser(ctx, name)
}
func (s *Service) UpdateUser(ctx context.Context, name string) (repository.User, error) {
	return s.Repository.UpdateUser(ctx, name)
}
func (s *Service) DeleteUser(ctx context.Context, name string) error {
	return s.Repository.DeleteUser(ctx, name)
}
func (s *Service) StoreOrder(ctx context.Context, order repository.Order) (repository.Order, error) {
	return s.Repository.StoreOrder(ctx, order)
}
func (s *Service) GetOrder(ctx context.Context, id int) (repository.Order, error) {
	return s.Repository.GetOrder(ctx, id)
}
func (s *Service) DeleteOrder(ctx context.Context, id int) error {
	return s.Repository.DeleteOrder(ctx, id)
}
func (s *Service) CreatePet(ctx context.Context, pet repository.Pet) (repository.Pet, error) {
	return s.Repository.CreatePet(ctx, pet)
}
func (s *Service) UpdatePet(ctx context.Context, pet repository.Pet) (repository.Pet, error) {
	return s.Repository.UpdatePet(ctx, pet)
}
func (s *Service) FindPetByStatus(ctx context.Context, status string) (repository.Pet, error) {
	return s.Repository.FindPetByStatus(ctx, status)
}
func (s *Service) FindPetById(ctx context.Context, id int) (repository.Pet, error) {
	return s.Repository.FindPetById(ctx, id)
}
func (s *Service) DeletePet(ctx context.Context, id int) error {
	return s.Repository.DeletePet(ctx, id)
}
