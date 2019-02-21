package form

type CaseCategory struct {
	RId int `json:"rid"`
	CId int `json:"cid"`
}

type Case struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	Price       int            `json:"price"`
	Type        int            `json:"type"`
	OwnerName   string         `json:"ownerName"`
	PhoneNumber string         `json:"phoneNumber"`
	Addr        string         `json:"addr"`
	Categorys   []CaseCategory `json:"category"`
}
