package model

type BlogArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (a BlogArticleTag) TableName() string {
	return "blog_article_tag"
}
