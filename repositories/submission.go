package repositories

import (
	"context"
	"database/sql"
	"github.com/BatuhanIlhan/gjg-casestudy/database/entities"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"time"
)

type SubmissionRepository struct {
	DB          *sql.DB
	IdGenerator func() string
	Clock       func() time.Time
}

type SubmissionCreatePayload struct {
	UserId string
	Score  float64
}

func NewSubmissionRepository(db *sql.DB) *SubmissionRepository {
	return &SubmissionRepository{db, uuid.NewString, time.Now}
}

func (r *SubmissionRepository) Create(ctx context.Context, payload SubmissionCreatePayload) (*entities.Submission, error) {

	now := r.Clock()
	entity := &entities.Submission{
		ID:        r.IdGenerator(),
		UserID:    payload.UserId,
		Score:     payload.Score,
		CreatedAt: now,
		UpdatedAt: now,
	}
	columns := []string{
		entities.SubmissionColumns.ID,
		entities.SubmissionColumns.UserID,
		entities.SubmissionColumns.Score,
		entities.SubmissionColumns.CreatedAt,
		entities.UserColumns.UpdatedAt,
	}

	if err := entity.Insert(ctx, r.DB, boil.Whitelist(columns...)); err != nil {
		return nil, err
	}
	return entity, nil
}
