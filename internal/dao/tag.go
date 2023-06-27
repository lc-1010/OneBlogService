package dao

import (
	"github.com/lc-1010/OneBlogService/internal/model"
	"github.com/lc-1010/OneBlogService/pkg/app"
)

func (d *Dao) GetTag(id uint32, state uint8) (model.BlogTag, error) {
	tag := model.BlogTag{Model: &model.Model{ID: id}, State: state}
	return tag.Get(d.engine)
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.BlogTag{State: state, Name: name}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.BlogTag, error) {
	tag := model.BlogTag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, crateBy string) error {
	tag := model.BlogTag{Name: name, State: state, Model: &model.Model{
		CreatedBy: crateBy,
	}}
	return tag.Create(d.engine)
}

func (d *Dao) CheckName(name string) (model.BlogTag, error) {

	tag := model.BlogTag{Name: name}
	return tag.CheckName(d.engine)
}

// UpdateTag  tag's function is db.Updates use values store key value
func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.BlogTag{Model: &model.Model{
		ID: id,
	}}
	values := map[string]any{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	tag.Update(d.engine, values)
	return nil
}
func (d *Dao) DeleteTag(id uint32) error {
	tag := model.BlogTag{Model: &model.Model{
		ID: id,
		//ModifiedBy: modifiedBy,
	}}
	return tag.Delete(d.engine)
}
