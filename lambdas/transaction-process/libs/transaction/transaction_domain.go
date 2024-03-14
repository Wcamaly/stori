package transaction

type TransactionRequest struct {
	UserID    string  `json:"userId"`
	Value     float64 `json:"value"`
	CreatedAt string  `json:"createdAt"`
}

// Read implements io.Reader.
func (t *TransactionRequest) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}

type TransactionResponse struct {
	ID        string  `json:"id"`
	UserID    string  `json:"userId"`
	Value     float64 `json:"value"`
	CreatedAt string  `json:"createdAt"`
}

func NewTransactionRequest(userId string, value float64, createdAt string) *TransactionRequest {
	return &TransactionRequest{
		UserID:    userId,
		Value:     value,
		CreatedAt: createdAt,
	}
}
