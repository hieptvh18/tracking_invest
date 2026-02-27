package services

import (
	"api_golang/models"
	"api_golang/repositories"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrEmailExists  = errors.New("email already exists")
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(req models.CreateUserRequest) (*models.User, error) {
	if _, ok := s.repo.GetByEmail(req.Email); ok {
		return nil, ErrEmailExists
	}
	u := &models.User{
		ID:    uuid.New().String(),
		Name:  req.Name,
		Age:   req.Age,
		Phone: req.Phone,
		Email: req.Email,
	}
	return s.repo.Create(u)
}

func (s *UserService) GetByID(id string) (*models.User, error) {
	u, ok := s.repo.GetByID(id)
	if !ok {
		return nil, ErrUserNotFound
	}
	return u, nil
}

func (s *UserService) List() []*models.User {
	return s.repo.List()
}

func (s *UserService) Update(id string, req models.UpdateUserRequest) (*models.User, error) {
	u, ok := s.repo.GetByID(id)
	if !ok {
		return nil, ErrUserNotFound
	}
	if req.Email != nil {
		existing, ok := s.repo.GetByEmail(*req.Email)
		if ok && existing.ID != id {
			return nil, ErrEmailExists
		}
		u.Email = *req.Email
	}
	if req.Name != nil {
		u.Name = *req.Name
	}
	if req.Age != nil {
		u.Age = *req.Age
	}
	if req.Phone != nil {
		u.Phone = *req.Phone
	}
	updated, ok := s.repo.Update(u)
	if !ok {
		return nil, ErrUserNotFound
	}
	return updated, nil
}

func (s *UserService) Delete(id string) error {
	if !s.repo.Delete(id) {
		return ErrUserNotFound
	}
	return nil
}
