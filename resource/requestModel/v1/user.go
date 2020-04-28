package v1

//ChangePassReqModel ...
type ChangePassReqModel struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
