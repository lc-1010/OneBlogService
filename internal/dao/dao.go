// dao 层进行了数据访问对象的封装，并针对业务所需的字段进行了处理。
package dao

import (
	"gorm.io/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
