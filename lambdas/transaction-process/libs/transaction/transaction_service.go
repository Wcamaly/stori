package transaction

import (
	"bytes"
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

func (t *TransactionConfig) CreateTransaction(payload *TransactionRequest) (*TransactionResponse, error) {
	payloadBytes, _ := json.Marshal(payload)

	data, err := http.Post(t.TransactionServiceURL+"/transaction", "application/json", bytes.NewBuffer(payloadBytes))

	if err != nil {
		println("Error to create transaction: ", err.Error())
		return nil, err
	}
	defer data.Body.Close()
	var res *TransactionResponse
	err = json.NewDecoder(data.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *TransactionConfig) CreateListTransaction(transactions []*TransactionRequest) ([]*TransactionResponse, error) {
	res := make([]*TransactionResponse, len(transactions))
	for i, transaction := range transactions {
		data, err := t.CreateTransaction(transaction)
		if err != nil {
			println("Error to create transactions: ", err.Error())
		}
		res[i] = data
	}
	return res, nil
}
