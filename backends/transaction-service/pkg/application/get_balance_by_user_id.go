package transaction

import (
	"context"
	"errors"
	"stori/transaction-service/pkg/domain/models"
	"stori/transaction-service/pkg/domain/transaction"
)

type GetBalanceByUserId struct {
	repository transaction.TransactionRepository
}

func NewGetBalanceByUserId(
	repository transaction.TransactionRepository,
) *GetBalanceByUserId {
	return &GetBalanceByUserId{
		repository: repository,
	}
}

type GetBalanceByUserIdDto struct {
	UserId models.ID `json:"userId"`
}

type GetBalanceByUserIdResponse struct {
	UserID  string  `json:"userId"`
	Balance float64 `json:"balance"`
}

func (tr *GetBalanceByUserId) Exec(ctx context.Context, payload *GetBalanceByUserIdDto) (*GetBalanceByUserIdResponse, error) {

	balance, err := tr.repository.FindBalanceByUserId(ctx, payload.UserId)
	if err != nil {
		return nil, errors.New("error getting balance")
	}
	return &GetBalanceByUserIdResponse{UserID: payload.UserId.String(), Balance: balance.Balance()}, nil
}
