package v1

//CustomerReq ...
type CustomerReq struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}
