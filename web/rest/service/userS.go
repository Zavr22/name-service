package service

import (
	"EMTestTask/pkg/model"
	"context"
	"github.com/google/uuid"
)

type User interface {
	CreateUser(ctx context.Context, user *model.User) (uuid.UUID, error)
	GetAllUsers(ctx context.Context, offset int) ([]*model.User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (model.User, error)
	UpdateProfile(ctx context.Context, userID uuid.UUID, input model.EnrichedFIO) error
	DeleteProfile(ctx context.Context, userID uuid.UUID) error
}

type UserService struct {
	userRepo User
}

func NewUserService(userRepo User) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.FIO) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) GetAllUsers(ctx context.Context, page int) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) GetUser(ctx context.Context, userID uuid.UUID) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) UpdateProfile(ctx context.Context, userID uuid.UUID, input model.EnrichedFIO) error {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) DeleteProfile(ctx context.Context, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}