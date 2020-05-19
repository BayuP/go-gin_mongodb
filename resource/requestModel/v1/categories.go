package v1

//CreateCatReq ...
type CreateCatReq struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
}

//UpdateCatReq ...
type UpdateCatReq struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}
