package form

type CaseCategory struct {
	RId          int    `json:"rid"`
	CategoryId   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	CaseId       int    `json:"caseId"`
}

type Case struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Type        int    `json:"type"`
	OwnerName   string `json:"ownerName"`
	PhoneNumber string `json:"phoneNumber"`
	Addr        string `json:"addr"`
}

type ComplexCaseCategory struct {
	CaseInfo     Case           `json:"caseInfo"`
	CategoryInfo []CaseCategory `json:"categoryInfo"`
}
