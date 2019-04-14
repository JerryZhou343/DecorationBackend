package form

//Category 分类信息
type Category struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Priority  int    `json:"priority"`
	ParentID  int    `json:"parentId"`
	Remark    string `json:"remark"`
	CreatedAt int    `json:"created_at"`
}

type PageInfo struct {
	ParentID int `json:"parentID"`
	Limit    int `json:"limit"`
	Offset   int `json:"offset"`
}
