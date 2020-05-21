package v1

//CreateProdReq ...
type CreateProdReq struct {
	Name       string `json:"name"`
	CategoryID string `json:"cat_id"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
}

//UpdateProdReq ...
type UpdateProdReq struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
}
