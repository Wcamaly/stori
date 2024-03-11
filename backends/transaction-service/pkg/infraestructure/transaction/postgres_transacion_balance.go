package transaction

import (
	_ "github.com/jackc/pgx/v5/stdlib"

	"stori/transaction-service/pkg/domain/models"
	"stori/transaction-service/pkg/domain/transaction"
)

type TransactionBalancePostgres struct {
	UserID  string  `db:"user_id"`
	Balance float64 `db:"balance"`
}

func (p *TransactionBalancePostgres) toDomain() (*transaction.TransactionBalance, error) {
	userID := models.ID(p.UserID)
	return transaction.NewTransactionBalance(userID, p.Balance), nil
}
