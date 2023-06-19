package modle

type BlogTag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (model BlogTag) TableName() string {
	return "blog_tag"
}
