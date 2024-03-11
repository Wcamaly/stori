package transaction

import (
	"context"
	"errors"
	"stori/transaction-service/pkg/domain/models"
	"stori/transaction-service/pkg/domain/transaction"
	"time"
)

type CreateTransaction struct {
	repository transaction.TransactionRepository
}

func NewCreateTransaction(
	repository transaction.TransactionRepository,
) *CreateTransaction {
	return &CreateTransaction{
		repository: repository,
	}
}

type CreateTransactionDto struct {
	UserID    models.ID `json:"userId"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateTransactionResponse struct {
	Id        string    `json:"id"`
	UserID    string    `json:"userId"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
}

func (tr *CreateTransaction) Exec(ctx context.Context, payload *CreateTransactionDto) (*CreateTransactionResponse, error) {

	id := models.GenerateUUID()
	newTransaction := transaction.NewTransaction(id, payload.UserID, payload.Value, payload.CreatedAt)
	err := tr.repository.Create(ctx, newTransaction)
	if err != nil {
		return nil, errors.New("error creating transaction")
	}
	return &CreateTransactionResponse{Id: id.String(), UserID: payload.UserID.String(), Value: payload.Value, CreatedAt: payload.CreatedAt}, nil
}
