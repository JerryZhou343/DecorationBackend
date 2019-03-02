package form

//Category 分类信息
type Category struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	ParentID int    `json:"parentId"`
	Remark   string `json:"remark"`
}
