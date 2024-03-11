package transaction

import (
	"context"
	"errors"
	"stori/transaction-service/pkg/domain/models"
	"stori/transaction-service/pkg/domain/transaction"
)

type GetMovementByUserId struct {
	repository transaction.TransactionRepository
}

func NewGetMovementByUserId(
	repository transaction.TransactionRepository,
) *GetMovementByUserId {
	return &GetMovementByUserId{
		repository: repository,
	}
}

type GetMovementByUserIdDto struct {
	UserId models.ID `json:"userId"`
}

type GetMovementByUserIdResponse struct {
	UserID    string `json:"userId"`
	Increment int    `json:"increment"`
	Decrement int    `json:"decrement"`
	Month     int    `json:"month"`
}

func (tr *GetMovementByUserId) Exec(ctx context.Context, payload *GetMovementByUserIdDto) ([]*GetMovementByUserIdResponse, error) {

	movements, err := tr.repository.FindMovementById(ctx, payload.UserId)
	if err != nil {
		return nil, errors.New("error getting movement")
	}
	response := make([]*GetMovementByUserIdResponse, len(movements))
	for i, movement := range movements {
		response[i] = &GetMovementByUserIdResponse{
			UserID:    payload.UserId.String(),
			Increment: movement.Increment(),
			Decrement: movement.Decrement(),
			Month:     int(movement.Month()),
		}
	}
	return response, nil

}
