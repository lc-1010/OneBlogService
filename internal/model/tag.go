package model

type BlogTag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (model BlogTag) TableName() string {
	return "blog_tag"
}

type UpdateTagRequest struct {
	Name      string `json:"name,omitempty" binding:"max=100"`              // 标签名称
	State     int    `json:"state,omitempty" binding:"oneof=0 1,default=1"` // 标签状态
	CreatedBy string `json:"created_by,omitempty" binding:"min=3,max=100"`  // 创建者
	TAGID     uint32 `json:"id,omitempty"`                                  // 文章id"
}
