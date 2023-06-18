package transformers

import (
	"github.com/BatuhanIlhan/gjg-casestudy/database/entities"
	"github.com/BatuhanIlhan/gjg-casestudy/models"
	"github.com/go-openapi/strfmt"
)

type UserTransformer func(entity *entities.User) *models.User

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

//func WalletList(wallets entities.WalletSlice) []*models.Wallet {
//	_users := make([]*models.Wallet, len(wallets))
//	for index, wallet := range wallets {
//		_users[index] = Wallet(wallet)
//	}
//	return _users
//}
