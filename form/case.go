package form

//CaseCategory 案例类型定义，和客户端交换数据模型
type CaseCategory struct {
	RID          int    `json:"RID"`
	CategoryID   int    `json:"categoryID"`
	CategoryName string `json:"categoryName"`
	CaseID       int    `json:"caseID"`
}

//Case 案例定义，和浏览器交换数据模型定义
type Case struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Type        int     `json:"type"`
	Priority    int     `json:"priority"`
	OwnerName   string  `json:"ownerName"`
	PhoneNumber string  `json:"phoneNumber"`
	Addr        string  `json:"addr"`
}

//ComplexCaseCategory 复合类型，包含案例信息和案例分类信息
type ComplexCaseCategory struct {
	CaseInfo     Case           `json:"caseInfo"`
	CategoryInfo []CaseCategory `json:"categoryInfo"`
}
