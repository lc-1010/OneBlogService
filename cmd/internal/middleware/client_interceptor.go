package middleware

import (
	"context"
	"time"

	"github.com/lc-1010/OneBlogService/global"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func defaultContextTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc
	if _, ok := ctx.Deadline(); !ok {
		defaultTimeout := 60 * time.Second
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

		ctx, span := global.Tracer.Tracer("client").Start(ctx, method)
		defer span.End()

		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}

		newCtx := metadata.NewOutgoingContext(ctx, md)

		return invoker(newCtx, method, req, resp, cc, opts...)
	}
}
