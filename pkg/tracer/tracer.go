package tracer

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

// NewJaegerTrancer set jaeger config
func NewJaegerTrancer(serviceNmae, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceNmae, // 服务名
		Sampler: &config.SamplerConfig{
			Type:  "const", // 采样类型
			Param: 1,       // 采样参数
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,            // 日志跟踪
			BufferFlushInterval: 1 * time.Second, // 缓冲区刷新间隔
			LocalAgentHostPort:  agentHostPort,   // 本地代理主机端口
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
