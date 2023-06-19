package transformers

import (
	"github.com/BatuhanIlhan/gjg-casestudy/database/entities"
	"github.com/BatuhanIlhan/gjg-casestudy/models"
	"github.com/go-openapi/strfmt"
)

type SubmissionTransformer func(entity *entities.Submission, newScore *float64) *models.Submission

func Submission(entity *entities.Submission, newScore *float64) *models.Submission {
	if entity == nil {
		return nil
	}

	return &models.Submission{
		ID:             strfmt.UUID(entity.ID),
		UserID:         strfmt.UUID(entity.UserID),
		SubmittedScore: entity.Score,
		NewScore:       *newScore,
		CreatedAt:      strfmt.DateTime(entity.CreatedAt),
		UpdatedAt:      strfmt.DateTime(entity.UpdatedAt),
	}
}

//func WalletList(wallets entities.WalletSlice) []*models.Wallet {
//	_users := make([]*models.Wallet, len(wallets))
//	for index, wallet := range wallets {
//		_users[index] = Wallet(wallet)
//	}
//	return _users
//}
