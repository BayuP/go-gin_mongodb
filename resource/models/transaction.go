package models

//Transactions ...
type Transactions struct {
	Base
	ID                 string               `json:"id"`
	CustomerID         string               `json:"customerID"`
	CashierID          string               `json:"cashierID"`
	DetailTransactions []DetailTransactions `json:"detailTransaction"`
	TotalPrice         int                  `json:"totalPrice"`
}

//DetailTransactions ...
type DetailTransactions struct {
	Base
	ID        string `json:"id"`
	ProductID string `json:"productId"`
	Price     int    `json:"price"`
}
