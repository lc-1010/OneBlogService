package tracer

import (
	"context"

	"github.com/lc-1010/OneBlogService/global"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

// BlogTrance is the tracer for blog
type BlogTrance struct {
	ServiceName string
}

func NewBlogTrace() gorm.Plugin {
	return BlogTrance{
		ServiceName: "blog",
	}
}

// Name is the name of this plugin
func (b BlogTrance) Name() string {
	return b.ServiceName
}

// Initialize is the initialization function of this plugin
// todo :增加其他事件处理
func (b BlogTrance) Initialize(db *gorm.DB) error {
	ctx := context.Context(context.Background())
	_, span := global.Tracer.Tracer("gorm").Start(ctx, "start-Initialize")
	defer span.End()
	// 注册 GORM 的回调钩子，用于创建和结束 OpenTelemetry 的 span
	// db.Callback().Create().Before("gorm:create").Register("ot:create", createCallback)
	//db.Callback().Update().Before("gorm:update").Register("ot:update", updateCallback)
	_ = db.Callback().Query().After("gorm:query").Register("ot:query", queryCallback)

	//db.Callback().Delete().Before("gorm:delete").Register("ot:delete", deleteCallback)
	return nil
}

func queryCallback(db *gorm.DB) {
	ctx := context.Context(context.Background())
	_, span := global.Tracer.Tracer("gorm").Start(ctx, "query",
		trace.WithAttributes(
			attribute.String("query", db.Statement.SQL.String()),
			attribute.String("table", db.Statement.Table),
		))

	defer span.End()
}
