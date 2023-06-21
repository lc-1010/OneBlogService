package model

import (
	"fmt"

	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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

func NewDBEngine(dbsetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dbsetting.UserName,
		dbsetting.Password,
		dbsetting.Host,
		dbsetting.DBName,
		dbsetting.Charset,
		dbsetting.ParseTime)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.Logger.LogMode(logger.Info)
	}

	DB, err := db.DB()
	if err != nil {
		return nil, err
	}

	DB.SetMaxIdleConns(dbsetting.MaxIdleConns)
	DB.SetMaxOpenConns(dbsetting.MaxOpenConns)
	return db, nil
}
