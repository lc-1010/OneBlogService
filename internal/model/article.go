package model

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

type CreateArticle struct {
	TagID         uint32 `json:"tag_id"`          // 标签ID
	Title         string `json:"title"`           // 文章标题"`
	Desc          string `json:"desc"`            // 文章简述"`
	CoverImageUrl string `json:"cover_image_url"` // 封面图片地址"`
	Content       string `json:"content"`         // 文章内容"`
	CreatedBy     string `json:"created_by"`      // 创建者"`
	State         uint8  `json:"state"`           // 状态"`
}

type UpdateArticle struct {
	TagID         uint32 `json:"tag_id,omitempty"` // 标签ID
	Title         string `json:"title,omitempty"`  // 文章标题"`
	Desc          string `json:"desc,omitempty"`   // 文章简述"`
	CoverImageUrl string `json:"cover_image_url"`  // 封面图片地址"`
	Content       string `json:"content"`          // 文章内容"`
	CreatedBy     string `json:"created_by"`       // 创建者"`
	State         uint8  `json:"state,omitempty"`  // 状态"`
	ModifiedBy    string `json:"modified_by"`      //"修改人"

}
