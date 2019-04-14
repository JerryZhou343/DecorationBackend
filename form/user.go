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

type UserInfo struct {
	Name   string `json:"name"`
	Roles  []int  `json:"roles"`
	Avatar string `json:"avatar"`
}
