package transformers

import (
	"github.com/BatuhanIlhan/gjg-casestudy/database/entities"
	"github.com/BatuhanIlhan/gjg-casestudy/models"
	"github.com/go-openapi/strfmt"
)

type SubmissionTransformer func(entity *entities.Submission) *models.Submission

//type UserListTransformer func(wallets entities.WalletSlice) []*models.Wallet

func Submission(entity *entities.Submission) *models.Submission {
	if entity == nil {
		return nil
	}

	return &models.Submission{
		ID:        strfmt.UUID(entity.ID),
		UserID:    strfmt.UUID(entity.UserID),
		Score:     entity.Score,
		CreatedAt: strfmt.DateTime(entity.CreatedAt),
		UpdatedAt: strfmt.DateTime(entity.UpdatedAt),
	}
}

//func WalletList(wallets entities.WalletSlice) []*models.Wallet {
//	_users := make([]*models.Wallet, len(wallets))
//	for index, wallet := range wallets {
//		_users[index] = Wallet(wallet)
//	}
//	return _users
//}
