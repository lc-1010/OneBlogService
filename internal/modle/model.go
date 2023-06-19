package modle

type Model struct {
	// id
	ID uint32 `gorm:"primary_key" json:"id"`
	// 创建时间
	CratedOn uint32 `json:"crated_on"`
	// 创建人
	CratedBy string `json:"crated_by"`
	// 修改时间
	ModifiedOn uint32 `json:"modified_on"`
	// 修改人
	ModifiedBy string `json:"modified_by"`
	// 删除时间
	DeletedOn uint32 `json:"deleted_on"`
	// 是否删除 0 未删 ， 1 已删
	IsDel uint8 `json:"is_del"`
}
