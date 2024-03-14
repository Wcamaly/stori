package transaction

import (
	_ "github.com/jackc/pgx/v5/stdlib"

	"stori/transaction-service/pkg/domain/models"
	"stori/transaction-service/pkg/domain/transaction"
)

type TransactionBalancePostgres struct {
	UserID  string  `db:"user_id"`
	Balance float64 `db:"balance"`
	Credit  float64 `db:"credit"`
	Debit   float64 `db:"debit"`
}

func (p *TransactionBalancePostgres) toDomain() (*transaction.TransactionBalance, error) {
	userID := models.ID(p.UserID)

	return transaction.NewTransactionBalance(userID, p.Balance, p.Credit, p.Debit), nil
}
