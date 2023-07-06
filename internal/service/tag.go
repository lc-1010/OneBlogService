package service

import (
	"github.com/lc-1010/OneBlogService/internal/model"
	"github.com/lc-1010/OneBlogService/pkg/app"
	"github.com/lc-1010/OneBlogService/pkg/errcode"
)

type ConutTagRequest struct {
	Name  string `form:"name,omitempty" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1" `
}

type GetTagListRequest struct {
	Name  string `form:"name,omitempty" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1" `
}

type CrateTagRquest struct {
	Name      string `form:"name,omitempty" binding:"max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1" `
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name,omitempty" binding:"max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1" `
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `uri:"id" binding:"required,gte=1"`
}

/********Service ********/

func (svc *Service) CountTag(param *ConutTagRequest) (int, error) {
	return svc.dao.CountTag(svc.ctx, param.Name, param.State)
}

func (svc *Service) GetTagList(param *GetTagListRequest, pager *app.Pager) ([]*model.BlogTag, error) {
	return svc.dao.GetTagList(svc.ctx, param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CrateTag(param *CrateTagRquest) error {
	find := svc.dao.CheckName(svc.ctx, param.Name)

	if find {
		return errcode.ErrorCrateTagExists
	}
	return svc.dao.CreateTag(svc.ctx, param.Name, param.State, param.CreatedBy)
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	del, err := svc.dao.DeleteTag(svc.ctx, param.ID)
	if err != nil {
		return err
	}
	if !del {
		return errcode.ErrorDeleteTagFail
	}
	return nil
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(svc.ctx, param.ID, param.Name, param.State, param.ModifiedBy)
}
