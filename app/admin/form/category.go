package form

type Category struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	ParentId int    `json:"parentId"`
	Remark   string `json:"remark"`
}
