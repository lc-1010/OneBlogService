package global

import (
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

var (
	// Tracer opentelemetry tracesdk
	Tracer *tracesdk.TracerProvider
)
