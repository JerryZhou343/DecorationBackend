package form

//Login 登录信息描述
type Login struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type NewPassWord struct {
	OldPassword string `json:"oldPassword"`
	NewPassWord string `json:"newPassword"`
}
