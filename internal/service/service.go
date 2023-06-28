package service

import (
	"context"

	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/internal/dao"
)

// Service
// 使用上下文和dao 的db engien 来处理具体的逻辑
type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
