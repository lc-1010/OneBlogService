package cmd_grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/lc-1010/OneBlogService/global"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func defaultContextTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc
	if _, ok := ctx.Deadline(); !ok {
		defaultTimeout := 3 * time.Second
		ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
	}
	return ctx, cancel
}

func StreamContextTimeout() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn,
		method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ctx, cancel := defaultContextTimeout(ctx)
		if cancel != nil {
			defer cancel()
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}

func UnaryContentTimeout() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string,
		req, resp any, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, cancel := defaultContextTimeout(ctx)
		if cancel != nil {
			defer cancel()
		}
		return invoker(ctx, method, req, resp, cc, opts...)
	}
}

func ClientTracing() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string,
		req, resp any, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		//上一级
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}
		spanOpts := []trace.SpanStartOption{
			trace.WithSpanKind(trace.SpanKindClient),
		}
		newCtx, span := global.Tracer.Tracer("client").
			Start(ctx, method, spanOpts...)
		defer span.End()

		for k, v := range md {
			span.SetAttributes(attribute.StringSlice(k, v))
			fmt.Println(k, v)
		}

		span.SetAttributes(attribute.StringSlice("service", []string{"flag-test-ok"}))

		newCtx = metadata.NewOutgoingContext(newCtx, md)
		fmt.Println(md, newCtx, "======")
		return invoker(newCtx, method, req, resp, cc, opts...)

	}
}
func UnaryInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		log.Printf("method=%s, request=%+v", method, req)
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			log.Printf("method=%s, error=%v", method, err)
		} else {
			log.Printf("method=%s, response=%+v", method, reply)
		}
		return err
	}

}
