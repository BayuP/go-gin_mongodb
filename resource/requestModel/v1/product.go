package v1

//CreateProdReq ...
type CreateProdReq struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

//UpdateProdReq ...
type UpdateProdReq struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
