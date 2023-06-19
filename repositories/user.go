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

type UserRepository struct {
	DB          *sql.DB
	IdGenerator func() string
	Clock       func() time.Time
}

type UserCreatePayload struct {
	DisplayName string
	CountryCode *string
	Points      *float64
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db, uuid.NewString, time.Now}
}

func (r *UserRepository) Create(ctx context.Context, payload UserCreatePayload) (*entities.User, error) {

	now := r.Clock()
	entity := &entities.User{
		ID:          r.IdGenerator(),
		Points:      null.Float64FromPtr(payload.Points),
		CountryCode: null.StringFromPtr(payload.CountryCode),
		DisplayName: payload.DisplayName,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	columns := []string{
		entities.UserColumns.ID,
		entities.UserColumns.Points,
		entities.UserColumns.CountryCode,
		entities.UserColumns.DisplayName,
		entities.UserColumns.CreatedAt,
		entities.UserColumns.UpdatedAt,
	}

	if err := entity.Insert(ctx, r.DB, boil.Whitelist(columns...)); err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *UserRepository) Get(ctx context.Context, id string) (*entities.UserWithRank, error) {
	user, err := entities.UserWithRanks(qm.Where("id = ?", id)).One(ctx, r.DB)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetLeaderBoard(ctx context.Context, limit, offset int, queries ...qm.QueryMod) (entities.UserWithRankSlice, error) {
	queries = append(queries, qm.Limit(limit), qm.Offset(offset))
	userWithRanks, err := entities.UserWithRanks(queries...).All(ctx, r.DB)
	if err != nil {
		return nil, err
	}
	return userWithRanks, nil
}
