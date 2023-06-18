package services

import (
	"context"
	"fmt"
	"github.com/BatuhanIlhan/gjg-casestudy/common/errors"
	"github.com/BatuhanIlhan/gjg-casestudy/database/entities"
	"github.com/BatuhanIlhan/gjg-casestudy/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

type UserCreatePayload struct {
	DisplayName string
	CountryCode *string
	Points      *float64
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) Create(ctx context.Context, payload UserCreatePayload) (*entities.User, *errors.ServiceError) {

	newUser, RepoErr := s.userRepo.Create(ctx, repositories.UserCreatePayload{
		DisplayName: payload.DisplayName,
		CountryCode: payload.CountryCode,
		Points:      payload.Points,
	})

	if RepoErr != nil {
		fmt.Println(RepoErr)
		return nil, errors.InternalServerError
	}
	return newUser, nil
}
