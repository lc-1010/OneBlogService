package service

type ArticleRequst struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequst struct {
	TagID uint32 `form:"tag_id" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	TagID         uint32 `form:"tag_id" binding:"gte=1"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	Content       string `form:"content" binding:"required,,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_img_url" binding:"required,url"`
	CreateBy      string `form:"create_by" binding:"required,min=2,max=100"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	TagID         uint32 `form:"tag_id" binding:"gte=1"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	Content       string `form:"content" binding:"required,,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_img_url" binding:"required,url"`
	ModifiedBy    string `form:"create_by" binding:"required,min=2,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type Article struct {
	ID            uint32 `json:"id"`
	Tag           uint32 `json:"tag"`
	State         uint8  `json:"state"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
}
