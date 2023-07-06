package dao

import (
	"context"

	"github.com/lc-1010/OneBlogService/internal/model"
	"github.com/lc-1010/OneBlogService/pkg/app"
)

// GetTag retrieves a BlogTag from the Dao using the provided id and state.
//
// Parameters:
//   - id: an unsigned 32-bit integer representing the id of the BlogTag.
//   - state: an unsigned 8-bit integer representing the state of the BlogTag.
//
// Returns:
//   - a BlogTag struct representing the retrieved BlogTag.
//   - an error, if any occurred.
func (d *Dao) GetTag(ctx context.Context, id uint32, state uint8) (model.BlogTag, error) {
	tag := model.BlogTag{Model: &model.Model{ID: id}, State: state}
	return tag.Get(ctx, d.engine)
}

// CountTag returns the count of tags with the given name and state.
//
// It takes the following parameters:
// - name: a string representing the name of the tag.
// - state: a uint8 representing the state of the tag.
//
// It returns an integer representing the count of tags and an error if any.
func (d *Dao) CountTag(ctx context.Context, name string, state uint8) (int, error) {
	tag := model.BlogTag{State: state, Name: name}
	return tag.Count(ctx, d.engine)
}

// GetTagList retrieves a list of BlogTag objects based on the given parameters.
//
// Parameters:
// - name: a string representing the name of the tag.
// - state: an unsigned 8-bit integer representing the state of the tag.
// - page: an integer representing the page number.
// - pageSize: an integer representing the number of items per page.
//
// Returns:
// - a slice of pointers to BlogTag objects.
// - an error if there was an issue retrieving the list.
func (d *Dao) GetTagList(ctx context.Context, name string, state uint8, page, pageSize int) ([]*model.BlogTag, error) {
	tag := model.BlogTag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(ctx, d.engine, pageOffset, pageSize)
}

// CreateTag creates a tag in the blog.
//
// name: the name of the tag.
// state: the state of the tag.
// crateBy: the user who created the tag.
// error: an error if the tag creation fails.
func (d *Dao) CreateTag(ctx context.Context, name string, state uint8, crateBy string) error {
	tag := model.BlogTag{Name: name, State: state, Model: &model.Model{
		CreatedBy: crateBy,
	}}
	return tag.Create(ctx, d.engine)
}

// CheckName checks the name of a BlogTag against the database
// and returns the corresponding BlogTag and an error if any.
// It takes a single parameter 'name' of type string and returns
// a BlogTag and an error.
func (d *Dao) CheckName(ctx context.Context, name string) bool {

	tag := model.BlogTag{Name: name}
	return tag.CheckName(ctx, d.engine)
}

// UpdateTag  tag's function is db.Updates use values store key value
func (d *Dao) UpdateTag(ctx context.Context, id uint32, name string, state uint8, modifiedBy string) error {
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

	err := tag.Update(ctx, d.engine, values)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTag deletes a tag with the given ID.
//
// It takes an ID of type uint32 as a parameter.
// It returns an error.
func (d *Dao) DeleteTag(ctx context.Context, id uint32) (bool, error) {
	tag := model.BlogTag{Model: &model.Model{
		ID: id,
		//ModifiedBy: modifiedBy,
	}}
	return tag.Delete(ctx, d.engine)

}
