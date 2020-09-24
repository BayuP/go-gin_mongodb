package v1

//TransactionRes ...
type TransactionRes struct {
	CustomerID           string                 `json:"customerId"`
	CashierID            string                 `json:"cashierId"`
	TransactionID        string                 `json:"transactionId"`
	TransactionDetailRes []TransactionDetailRes `json:"transactionDetailRes"`
}

//TransactionDetailRes ...
type TransactionDetailRes struct {
	ID            string `json:"id"`
	TranscationID string `json:"transactionId"`
	ProductID     string `json:"productId"`
	Price         int    `json:"price"`
}
