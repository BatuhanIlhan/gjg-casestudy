package services

import (
	"context"
	"fmt"
	"github.com/BatuhanIlhan/gjg-casestudy/common/errors"
	"github.com/BatuhanIlhan/gjg-casestudy/database/entities"
	"github.com/BatuhanIlhan/gjg-casestudy/repositories"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func (s *UserService) Get(ctx context.Context, id string) (*entities.UserWithRank, *errors.ServiceError) {
	user, RepoErr := s.userRepo.Get(ctx, id)
	if RepoErr != nil {
		fmt.Println(RepoErr)
		return nil, errors.InternalServerError
	}
	return user, nil
}

func (s *UserService) GetLeaderBoard(ctx context.Context, limit, offset int, queries ...qm.QueryMod) (entities.UserWithRankSlice, *errors.ServiceError) {
	userWithRanks, RepoErr := s.userRepo.GetLeaderBoard(ctx, limit, offset, queries...)
	if RepoErr != nil {
		fmt.Println(RepoErr)
		return nil, errors.InternalServerError
	}
	return userWithRanks, nil
}
