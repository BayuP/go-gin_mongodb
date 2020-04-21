package v1

//LoginResponse ...
type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}
