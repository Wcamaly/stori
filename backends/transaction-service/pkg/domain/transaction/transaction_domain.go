package transaction

import (
	"context"
	"stori/transaction-service/pkg/config/errors"
	"stori/transaction-service/pkg/domain/models"
	"time"
)

var (
	ErrorTransactionInternal = errors.Define("transaction.internal_error")
)

type TransactionRepository interface {
	Count(ctx context.Context, filter *TransactionFilter) (int, error)
	Find(ctx context.Context, filter *TransactionFilter) ([]*Transaction, error)
	FindByUserId(ctx context.Context, userId models.ID) ([]*Transaction, error)
	FindById(ctx context.Context, id models.ID) (*Transaction, error)
	Create(ctx context.Context, c *Transaction) error
	Update(ctx context.Context, c *Transaction) error
	FindBalanceByUserId(ctx context.Context, userId models.ID) (*TransactionBalance, error)
	FindBalance(ctx context.Context) ([]*TransactionBalance, error)
	FindMovementById(ctx context.Context, userId models.ID) ([]*TransactionMovement, error)
}

type TransactionBalance struct {
	userId  models.ID
	balance float64
	debit   float64
	credit  float64
}

type TransactionMovement struct {
	userId    models.ID
	increment int
	decrement int
	month     time.Month
}

type Transaction struct {
	ID        models.ID
	userID    models.ID
	value     float64
	createdAt time.Time
}

func NewTransactionBalance(userId models.ID, balance float64) *TransactionBalance {
	return &TransactionBalance{userId: userId, balance: balance}
}

func (u *TransactionBalance) UserId() models.ID {
	return u.userId
}

func (u *TransactionBalance) Balance() float64 {
	return u.balance
}

func (u *TransactionBalance) Debit() float64 {
	return u.debit
}

func (u *TransactionBalance) Credit() float64 {
	return u.credit
}

func NewTransactionMovement(userId models.ID, increment int, decrement int, month time.Month) *TransactionMovement {
	return &TransactionMovement{userId: userId, increment: increment, decrement: decrement, month: month}
}

func (u *TransactionMovement) UserId() models.ID {
	return u.userId
}
func (u *TransactionMovement) Increment() int {
	return u.increment
}
func (u *TransactionMovement) Decrement() int {
	return u.decrement
}
func (u *TransactionMovement) Month() time.Month {
	return u.month
}

func NewTransaction(id models.ID, userId models.ID, value float64, createdAt time.Time) *Transaction {
	return &Transaction{
		ID: id, userID: userId, value: value, createdAt: createdAt,
	}
}

func (u *Transaction) Id() models.ID {
	return u.ID
}

func (u *Transaction) UserId() models.ID {
	return u.userID
}

func (u *Transaction) Value() float64 {
	return u.value
}

func (u *Transaction) CreatedAt() time.Time {
	return u.createdAt
}
