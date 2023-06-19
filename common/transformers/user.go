package transformers

import (
	"github.com/BatuhanIlhan/gjg-casestudy/database/entities"
	"github.com/BatuhanIlhan/gjg-casestudy/models"
	"github.com/go-openapi/strfmt"
)

type UserTransformer func(entity *entities.User) *models.User
type UserWithRankTransformer func(entity *entities.UserWithRank) *models.User
type UserWithRankListTransformer func(entity entities.UserWithRankSlice) []*models.User

//type UserListTransformer func(wallets entities.WalletSlice) []*models.Wallet

func User(entity *entities.User) *models.User {
	if entity == nil {
		return nil
	}
	countryCode := &entity.CountryCode.String
	return &models.User{
		ID:          strfmt.UUID(entity.ID),
		CountryCode: countryCode,
		Points:      entity.Points.Float64,
		CreatedAt:   strfmt.DateTime(entity.CreatedAt),
		UpdatedAt:   strfmt.DateTime(entity.UpdatedAt),
	}
}

func UserWithRank(entity *entities.UserWithRank) *models.User {
	if entity == nil {
		return nil
	}
	countryCode := &entity.CountryCode.String
	return &models.User{
		ID:          strfmt.UUID(entity.ID.String),
		CountryCode: countryCode,
		Points:      entity.Points.Float64,
		CreatedAt:   strfmt.DateTime(entity.CreatedAt.Time),
		UpdatedAt:   strfmt.DateTime(entity.UpdatedAt.Time),
		Rank:        entity.Rank.Int64,
	}
}

func UserWithRankList(entity entities.UserWithRankSlice) []*models.User {
	if entity == nil {
		return nil
	}
	_leaderboard := make([]*models.User, len(entity))
	for index, user := range entity {
		_leaderboard[index] = UserWithRank(user)
	}
	return _leaderboard
}
