package services

import (
	"context"
	"github.com/BatuhanIlhan/gjg-casestudy/common/errors"
	"github.com/BatuhanIlhan/gjg-casestudy/database/entities"
	"github.com/BatuhanIlhan/gjg-casestudy/repositories"
)

type SubmissionService struct {
	submissionRepo *repositories.SubmissionRepository
}

type SubmissionCreatePayload struct {
	UserId string
	Score  float64
}

func NewSubmissionService(repo *repositories.SubmissionRepository) *SubmissionService {
	return &SubmissionService{submissionRepo: repo}
}

func (s *SubmissionService) Create(ctx context.Context, payload SubmissionCreatePayload) (*entities.Submission, *float64, *errors.ServiceError) {

	newSubmission, newScore, RepoErr := s.submissionRepo.Create(ctx, repositories.SubmissionCreatePayload{
		UserId: payload.UserId,
		Score:  payload.Score,
	})

	if RepoErr != nil {
		return nil, nil, errors.InternalServerError
	}
	return newSubmission, newScore, nil
}
