package tracer

import (
	"log"
	"os"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// NewJaegerTrancer set jaeger config
func NewJaegerTrancer(serviceName, agentHost, agentPort string) (*tracesdk.TracerProvider, error) {
	//打印本地测试
	logger := log.New(os.Stdout, "jaeger: ", log.LstdFlags)
	//exporter, err := stdout.New(stdout.WithPrettyPrint())

	exp, err := jaeger.New(jaeger.WithAgentEndpoint( // 参考 https://github.com/owncloud/ocis/blob/a8ff963166ecd9adf3f44aa6fa9fe68f53517d05/ocis-pkg/tracing/tracing.go#L65
		jaeger.WithAgentHost(agentHost),
		jaeger.WithAgentPort(agentPort),
		jaeger.WithLogger(logger),
	))

	if err != nil {
		return nil, err
	}
	//_ = exp

	tp := tracesdk.NewTracerProvider(
		// 使用给定的批处理器配置追踪器提供程序
		tracesdk.WithBatcher(exp, tracesdk.WithMaxExportBatchSize(10)),

		// 使用给定的资源配置追踪器提供程序
		tracesdk.WithResource(resource.NewWithAttributes(
			// 资源的模式 URL，用于标识资源的模式
			semconv.SchemaURL,

			// 设置服务名称属性，将服务名称与追踪数据相关联
			semconv.ServiceName(serviceName),

			// 在资源中添加环境属性，用于标识当前的环境
			attribute.String("env", "dev"),

			// 在资源中添加实例属性，用于标识当前的实例编号
			attribute.Int64("instance", 1),
		),
		),
	)

	return tp, nil
}
