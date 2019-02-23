package form

type PicCategory struct {
	PicId        int    `json:"picId"`
	CategoryId   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

type Picture struct {
	PicId  int    `json:"picId"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Addr   string `json:"addr"`
}

type ComplexPicCategory struct {
	PicInfo      Picture     `json:"picture"`
	CategoryInfo PicCategory `json:"categoryInfo"`
}
