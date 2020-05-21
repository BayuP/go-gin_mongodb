package v1

//ProductResponse ...
type ProductResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
}
