package transaction

import (
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"stori/transaction-service/pkg/domain/models"
	"stori/transaction-service/pkg/domain/transaction"
)

type TransactionMovementPostgres struct {
	UserID     string `db:"user_id"`
	Increments int    `db:"increments"`
	Decrements int    `db:"decrements"`
	Month      int    `db:"month"`
}

func (p *TransactionMovementPostgres) toDomain() (*transaction.TransactionMovement, error) {
	userID := models.ID(p.UserID)
	month := time.Month(p.Month)
	return transaction.NewTransactionMovement(userID, p.Increments, p.Decrements, month), nil
}
