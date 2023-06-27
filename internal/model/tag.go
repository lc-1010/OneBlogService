package model

import (
	"github.com/lc-1010/OneBlogService/pkg/app"
	"gorm.io/gorm"
)

type BlogTag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}
type UpdateTagRequest struct {
	Name      string `json:"name,omitempty" binding:"max=100"`              // 标签名称
	State     int    `json:"state,omitempty" binding:"oneof=0 1,default=1"` // 标签状态
	CreatedBy string `json:"created_by,omitempty" binding:"min=3,max=100"`  // 创建者
	TAGID     uint32 `json:"id,omitempty"`                                  // 文章id"
}

type TagSwagger struct {
	List  []*BlogTag
	Pager *app.Pager
}

func (model BlogTag) TableName() string {
	return "blog_tag"
}

func NewBlogTag() *BlogTag {
	return &BlogTag{
		Model: &Model{},
	}
}

func (t BlogTag) Count(db *gorm.DB) (int, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (t BlogTag) List(db *gorm.DB, pageOffset, pageSize int) ([]*BlogTag, error) {
	var tags []*BlogTag
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, err
}

func (t BlogTag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t BlogTag) CheckName(db *gorm.DB) (BlogTag, error) {
	var tag BlogTag
	err := db.Where("name = ?", t.Name).First(&tag).Error
	if err != nil {
		return tag, err
	}
	return tag, nil
}

// Update Update  BlogTag
// When using the Update method to update data,
// it is important to ensure that the structure
// passed as a parameter is not nil to avoid null
// pointer exceptions. Therefore, it is necessary to
// initialize the structure before using it to avoid null pointer exceptions.
func (t BlogTag) Update(db *gorm.DB, value any) error {
	tag := NewBlogTag()
	return db.Model(tag).Where("id = ? AND is_del = ?", t.ID, 0).Updates(value).Error
}

func (t BlogTag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}

func (t BlogTag) Get(db *gorm.DB) (BlogTag, error) {
	var tag BlogTag
	err := db.Where("id = ? and is_del = ? and state = ?",
		t.ID, t.IsDel, t.State).First(&tag).Error
	if err != nil {
		return tag, err
	}
	return tag, nil
}
