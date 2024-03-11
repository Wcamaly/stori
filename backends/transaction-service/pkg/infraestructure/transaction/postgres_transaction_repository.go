package transaction

import (
	"context"
	"stori/transaction-service/pkg/config/errors"
	"stori/transaction-service/pkg/domain/models"
	"stori/transaction-service/pkg/domain/transaction"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var _ transaction.TransactionRepository = (*TransactionPostgressRepository)(nil)

func NewPostgresTransactionRepository(db *sqlx.DB) *TransactionPostgressRepository {
	return &TransactionPostgressRepository{
		db: db,
	}
}

type TransactionPostgressRepository struct {
	db *sqlx.DB
}

func (u TransactionPostgressRepository) Count(ctx context.Context, filter *transaction.TransactionFilter) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (u TransactionPostgressRepository) Find(ctx context.Context, filter *transaction.TransactionFilter) ([]*transaction.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (u TransactionPostgressRepository) FindByUserId(ctx context.Context, userId models.ID) ([]*transaction.Transaction, error) {

	var transactioPostgres []*TransactionPostgres
	if err := u.db.SelectContext(
		ctx,
		&transactioPostgres,
		`SELECT distinct id, t.user_id, t.created_at, t.value FROM public.transaction AS t WHERE (t.user_id = $1)`,
		userId,
	); err != nil {
		return nil, errors.Wrap(
			transaction.ErrorTransactionInternal,
			err,
			"contract internal error",
			errors.WithMetadata("userId", userId),
		)
	}

	transactios := make([]*transaction.Transaction, len(transactioPostgres))
	var err error
	for i, c := range transactioPostgres {
		transactios[i], err = c.toDomain()
		if err != nil {
			return nil, err
		}
	}
	return transactios, nil
}

func (u TransactionPostgressRepository) FindById(ctx context.Context, id models.ID) (*transaction.Transaction, error) {

	var transactioPostgres TransactionPostgres
	if err := u.db.GetContext(
		ctx,
		&transactioPostgres,
		`SELECT distinct id, t.user_id, t.created_at, t.value FROM public.transaction AS t WHERE (t.id = $1)`,
		id,
	); err != nil {
		return nil, errors.Wrap(
			transaction.ErrorTransactionInternal,
			err,
			"contract internal error",
			errors.WithMetadata("id", id),
			//errors.WithMetadata("query", sqlQuery),
			//errors.WithMetadata("args", args),
		)
	}

	return transactioPostgres.toDomain()
}

func (u TransactionPostgressRepository) Create(ctx context.Context, us *transaction.Transaction) error {
	tx, err := u.db.BeginTxx(ctx, nil)
	if err != nil {
		return errors.New(transaction.ErrorTransactionInternal, "error beginning transaction Transaction")
	}
	_, err = tx.NamedExecContext(ctx, `
		INSERT INTO public.transaction (id, user_id, value, created_at)
		VALUES (:id, :user_id, :value, :created_at)`,
		map[string]interface{}{
			"id":         us.Id(),
			"user_id":    us.UserId(),
			"value":      us.Value(),
			"created_at": us.CreatedAt(),
		})
	if err != nil {
		println("err: ", err.Error())
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (u TransactionPostgressRepository) Update(ctx context.Context, c *transaction.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (u TransactionPostgressRepository) FindBalanceByUserId(ctx context.Context, userId models.ID) (*transaction.TransactionBalance, error) {
	var transactionBalancePostgres TransactionBalancePostgres
	if err := u.db.GetContext(
		ctx,
		&transactionBalancePostgres,
		`SELECT t.user_id, sum(t.value) as balance from public.transaction as t WHERE (t.user_id = $1) GROUP BY user_id `,
		userId,
	); err != nil {
		return nil, errors.Wrap(
			transaction.ErrorTransactionInternal,
			err,
			"contract internal error",
			errors.WithMetadata("userId", userId),
		)
	}
	return transactionBalancePostgres.toDomain()
}

func (u TransactionPostgressRepository) FindBalance(ctx context.Context) ([]*transaction.TransactionBalance, error) {
	var transactionsBalancePostgres []*TransactionBalancePostgres
	if err := u.db.SelectContext(
		ctx,
		&transactionsBalancePostgres,
		`SELECT t.user_id, sum(t.value) as balance from public.transaction as t GROUP BY user_id`,
	); err != nil {
		return nil, errors.Wrap(
			transaction.ErrorTransactionInternal,
			err,
			"contract internal error",
		)
	}

	println("transactionsBalancePostgres", transactionsBalancePostgres)

	transactions := make([]*transaction.TransactionBalance, len(transactionsBalancePostgres))
	var err error
	for i, c := range transactionsBalancePostgres {
		transactions[i], err = c.toDomain()
		if err != nil {
			println("err", err.Error())
			return nil, err
		}
	}
	return transactions, nil
}

func (u TransactionPostgressRepository) FindMovementById(ctx context.Context, userId models.ID) ([]*transaction.TransactionMovement, error) {
	var transactionsMovementPostgres []*TransactionMovementPostgres
	if err := u.db.SelectContext(
		ctx,
		&transactionsMovementPostgres,
		`SELECT
					user_id,
					COUNT(*) FILTER (WHERE value > 0) AS increments,
					COUNT(*) FILTER (WHERE value < 0) AS decrements,
					DATE_PART('month', created_at) AS month
			FROM
					transaction 
			WHERE (user_id = $1)
			GROUP BY
					user_id,
					DATE_PART('month', created_at)
			ORDER BY
					user_id,
					month `,
		userId,
	); err != nil {
		println("err", err.Error())
		return nil, errors.Wrap(
			transaction.ErrorTransactionInternal,
			err,
			"contract internal error",
			errors.WithMetadata("userId", userId),
		)
	}
	movements := make([]*transaction.TransactionMovement, len(transactionsMovementPostgres))
	var err error
	for i, c := range transactionsMovementPostgres {
		movements[i], err = c.toDomain()
		if err != nil {
			return nil, err
		}
	}
	return movements, nil
}
