package transaction

import (
	_ "github.com/jackc/pgx/v5/stdlib"

	"stori/transaction-service/pkg/config/errors"
	"stori/transaction-service/pkg/domain/models"
	"stori/transaction-service/pkg/domain/transaction"
	"time"
)

var (
	ErrDomainConversion = errors.Define("transactiondto.toDomain_error")
)

type TransactionDto struct {
	ID        string  `json:"id"`
	UserID    string  `json:"userId"`
	Value     float64 `json:"value"`
	CreatedAt string  `json:"createdAt"`
}

type TransactionPostgres struct {
	ID        string  `db:"id"`
	UserID    string  `db:"user_id"`
	Value     float64 `db:"value"`
	CreatedAt string  `db:"created_at"`
}

func (p *TransactionPostgres) toDomain() (*transaction.Transaction, error) {
	id, err := models.NewID(p.ID)
	if err != nil {
		return nil, err
	}
	userID := models.ID(p.UserID)                           // Convert p.UserID to ID type
	createdAt, err := time.Parse(time.RFC3339, p.CreatedAt) // Convert p.CreatedAt to time.Time
	if err != nil {
		return nil, err
	}
	return transaction.NewTransaction(id, userID, p.Value, createdAt), nil
}

func (u TransactionDto) toDomain() (*transaction.Transaction, error) {
	id, err := models.NewID(u.ID)
	if err != nil {
		return nil, err
	}

	UserId, err := models.NewRequiredString(u.UserID)
	if err != nil {
		return nil, errors.New(ErrDomainConversion, "missing userID")
	}

	CreatedAt, err := models.NewRequiredTime(u.CreatedAt)
	if err != nil {
		return nil, errors.New(ErrDomainConversion, "missing createdAt")
	}

	return transaction.NewTransaction(id, models.ID(UserId), u.Value, time.Time(*CreatedAt)), nil
}

func fromDomain(u *transaction.Transaction) TransactionDto {
	return TransactionDto{
		ID:        string(u.Id()),
		UserID:    string(u.UserId()),
		Value:     u.Value(),
		CreatedAt: u.CreatedAt().Format(time.RFC3339),
	}
}
