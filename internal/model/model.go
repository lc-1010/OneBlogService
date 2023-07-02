package model

import (
	"fmt"
	"time"

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
	CreatedOn uint32 `json:"created_on"`
	// 创建人
	CreatedBy string `json:"created_by"`
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

/**************use hook *******/
// 初始化在update 时 &BlogTag{ Mode :&model{}}

func (t *Model) BeforeUpdate(db *gorm.DB) error {
	nowTime := time.Now().Unix()

	t.ModifiedOn = uint32(nowTime)

	return nil
}

func (t *Model) BeforeCreate(db *gorm.DB) error {
	nowTime := time.Now().Unix()

	t.CreatedOn = uint32(nowTime)

	return nil
}

func (t *Model) BeforeDelete(db *gorm.DB) error {
	nowTime := time.Now().Unix()

	t.DeletedOn = uint32(nowTime)

	return nil
}
