package transaction

import (
	"context"
	"errors"
	"stori/transaction-service/pkg/domain/transaction"
)

type GetBalance struct {
	repository transaction.TransactionRepository
}

func NewGetBalance(
	repository transaction.TransactionRepository,
) *GetBalance {
	return &GetBalance{
		repository: repository,
	}
}

type GetBalanceResponse struct {
	UserID  string  `json:"userId"`
	Balance float64 `json:"balance"`
}

func (tr *GetBalance) Exec(ctx context.Context) ([]*GetBalanceResponse, error) {

	balances, err := tr.repository.FindBalance(ctx)
	if err != nil {
		return nil, errors.New("error getting balance")
	}

	response := make([]*GetBalanceResponse, len(balances))
	for i, balance := range balances {
		response[i] = &GetBalanceResponse{UserID: balance.UserId().String(), Balance: balance.Balance()}
	}
	return response, nil
}
