package transaction

import (
	"context"
	"errors"
	"stori/transaction-service/pkg/domain/models"
	"stori/transaction-service/pkg/domain/transaction"

	"time"
)

type GetTransactionById struct {
	repository transaction.TransactionRepository
}

func NewGetTransactionById(
	repository transaction.TransactionRepository,
) *GetTransactionById {
	return &GetTransactionById{
		repository: repository,
	}
}

type GetTransactionByIdDto struct {
	ID models.ID `json:"Id"`
}

type GetTransactionByIdResponse struct {
	Id        string    `json:"id"`
	UserID    string    `json:"userId"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
}

func (tr *GetTransactionById) Exec(ctx context.Context, payload *GetTransactionByIdDto) (*GetTransactionByIdResponse, error) {

	tran, err := tr.repository.FindById(ctx, payload.ID)
	if err != nil {
		return nil, errors.New("error getting transaction")
	}
	return &GetTransactionByIdResponse{Id: tran.Id().String(), UserID: tran.UserId().String(), Value: tran.Value(), CreatedAt: tran.CreatedAt()}, nil
}
