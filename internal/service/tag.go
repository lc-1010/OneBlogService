package service

type ConutTagRequest struct {
	Name  string `form:"name,omitempty" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1" `
}

type TagListRequest struct {
	Name  string `form:"name,omitempty" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1" `
}

type CrateTagRquest struct {
	Name      string `form:"name,omitempty" binding:"max=100"`
	CreatedBy string `form:"crated_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1" `
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name,omitempty" binding:"max=100"`
	CreatedBy  string `form:"crated_by" binding:"required,min=3,max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1" `
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
