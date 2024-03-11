package transaction

import (
	"context"
	"errors"
	"stori/transaction-service/pkg/domain/models"
	"stori/transaction-service/pkg/domain/transaction"

	"time"
)

type GetTransactionByUserId struct {
	repository transaction.TransactionRepository
}

func NewGetTransactionByUserId(
	repository transaction.TransactionRepository,
) *GetTransactionByUserId {
	return &GetTransactionByUserId{
		repository: repository,
	}
}

type GetTransactionByUserIdDto struct {
	ID models.ID `json:"Id"`
}

type GetTransactionByUserIdResponse struct {
	Id        string    `json:"id"`
	UserID    string    `json:"userId"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
}

func (tr *GetTransactionByUserId) Exec(ctx context.Context, payload *GetTransactionByUserIdDto) ([]*GetTransactionByUserIdResponse, error) {

	tran, err := tr.repository.FindByUserId(ctx, payload.ID)
	if err != nil {
		return nil, errors.New("error getting transaction")
	}

	var response []*GetTransactionByUserIdResponse
	for _, t := range tran {
		response = append(response, &GetTransactionByUserIdResponse{Id: t.Id().String(), UserID: t.UserId().String(), Value: t.Value(), CreatedAt: t.CreatedAt()})
	}
	return response, nil
}
