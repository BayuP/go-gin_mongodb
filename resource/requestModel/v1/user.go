package v1

//ChangePassReqModel ...
type ChangePassReqModel struct {
	ID          string `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

//CreateUserReq ...
type CreateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   string `json:"roleId"`
}

//EditUserReq ...
type EditUserReq struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	RoleID   string `json:"RoleId"`
}
