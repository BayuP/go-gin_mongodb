package v1

//TransactionReq ...
type TransactionReq struct {
	CustomerID           string                 `json:"customerId"`
	TransactionDetailReq []TransactionDetailReq `json:"transactionDetail"`
}

//TransactionDetailReq ...
type TransactionDetailReq struct {
	ProductID string `json:"productId"`
	Price     int    `json:"price"`
}
