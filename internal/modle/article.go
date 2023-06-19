package modle

type BlogArticle struct {
	*Model
	Title string `json:"title"`
	// 文章简述
	Desc string `json:"desc"`
	// 封面图片地址
	CoverImageUrl string `json:"cover_image_url"`
	// 状态 0 禁用，1 启用
	State uint8 `json:"state"`
}

func (model BlogArticle) TableName() string {
	return "blog_article"
}
