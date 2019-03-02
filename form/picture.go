package form

//PicCategory 图片分类信息
type PicCategory struct {
	PicID        int    `json:"picID"`
	CategoryID   int    `json:"categoryID"`
	CategoryName string `json:"categoryName"`
}

//Picture 图片信息描述
type Picture struct {
	PicID  int    `json:"picID"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Addr   string `json:"addr"`
}

//ComplexPicCategory 复合图片信息，包含图片信息和图片分类信息
type ComplexPicCategory struct {
	PicInfo      Picture     `json:"picture"`
	CategoryInfo PicCategory `json:"categoryInfo"`
}
