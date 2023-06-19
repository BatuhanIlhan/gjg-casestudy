package repositories

import (
	"context"
	"database/sql"
	"github.com/BatuhanIlhan/gjg-casestudy/database/entities"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func (r *SubmissionRepository) Create(ctx context.Context, payload SubmissionCreatePayload) (*entities.Submission, *float64, error) {
	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		return nil, nil, err
	}

	defer func() { _ = tx.Rollback() }()

	// check if user exist
	user, err := entities.Users(qm.Where("id = ?", payload.UserId), qm.For("update")).One(ctx, tx)
	if err != nil {
		return nil, nil, err
	}

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
	newScore := user.Points.Float64 + payload.Score
	user.Points = null.Float64From(newScore)
	user.UpdatedAt = now
	updateFields := []string{entities.UserColumns.Points, entities.UserColumns.UpdatedAt}
	_, err = user.Update(ctx, tx, boil.Whitelist(updateFields...))
	if err != nil {
		return nil, nil, err
	}
	if err := entity.Insert(ctx, tx, boil.Whitelist(columns...)); err != nil {
		return nil, nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, nil, err
	}

	return entity, &newScore, nil
}
