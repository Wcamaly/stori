package config

import (
	"stori/transaction-service/pkg/domain/transaction"
	trct "stori/transaction-service/pkg/infraestructure/transaction"
	"stori/transaction-service/pkg/libs/sql"
	_ "stori/transaction-service/pkg/libs/sql"
)

type Dependencies struct {
	TransactionRepository transaction.TransactionRepository
}

func BuildDependencies(config *Config) (*Dependencies, error) {
	db, err := sql.NewDB(config.DBConfig)

	if err != nil {
		return nil, err
	}

	TransactionRepository := trct.NewPostgresTransactionRepository(db)

	return &Dependencies{
		TransactionRepository,
	}, nil
}
