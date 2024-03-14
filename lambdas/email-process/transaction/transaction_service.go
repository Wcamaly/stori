package transaction

import (
	"encoding/json"
	"net/http"
)

type TransactionConfig struct {
	TransactionServiceURL string `json:"transactionServiceURL"`
	Timeout               int    `json:"timeout"`
}

func NewTransactionService(transactionServiceURL string, timeout int) *TransactionConfig {
	return &TransactionConfig{
		TransactionServiceURL: transactionServiceURL,
		Timeout:               timeout,
	}
}

func (t *TransactionConfig) GetBalanceByUserId(userId string) (*BalanceResponse, error) {
	data, err := http.Get(t.TransactionServiceURL + "/transaction/balance/" + userId)
	if err != nil {
		return nil, err
	}
	defer data.Body.Close()

	var res *BalanceResponse
	err = json.NewDecoder(data.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *TransactionConfig) GetMovementByUserId(userId string) ([]*MovementResponse, error) {
	data, err := http.Get(t.TransactionServiceURL + "/transaction/movement/" + userId)
	if err != nil {
		return nil, err
	}
	defer data.Body.Close()

	var res []*MovementResponse
	err = json.NewDecoder(data.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
