package middleware

import (
	"context"
	"log"
	"runtime/debug"
	"time"

	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/pkg/errcode"
	"github.com/lc-1010/OneBlogService/pkg/metatext"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Recovery(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	defer func() {
		if e := recover(); e != nil {
			recoverLog := "recovery log: method:%s, message:%v, stack:%s"
			log.Printf(recoverLog, info.FullMethod,
				e, string(debug.Stack()[:]))
		}
	}()
	return handler(ctx, req)
}

// ErrorLog logs any errors that occur during the execution of the grpc server handler function.
//
// ErrorLog takes in the following parameters:
//   - ctx: the context.Context object that represents the request context.
//   - req: the input request object of any type.
//   - info: the *grpc.UnaryServerInfo object that contains information about the server and method being called.
//   - handler: the grpc.UnaryHandler function that handles the request and returns a response and an error.
//
// ErrorLog returns the following values:
//   - resp: the response object of any type returned by the handler function.
//   - err: an error object that contains any error that occurred during the execution of the handler function.
func ErrorLog(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		errLog := "error log:method:%s,code :%v,messsage:%v,details:%v"
		s := errcode.FromError(err)
		log.Printf(errLog, info.FullMethod, s.Code(), s.Err())
	}
	return resp, err
}

func AccessLog(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	requestLog := "access request log:method :%s,begin_time:%d,request:%v"
	beginTime := time.Now().Local().Unix()
	log.Printf(requestLog, info.FullMethod, beginTime, req)

	//handler process
	resp, err := handler(ctx, req)

	responseLog := "access response log:method:%s,begin_time:%d,end_time:%d response:%v"
	endTime := time.Now().Local().Unix()
	log.Printf(responseLog, info.FullMethod, beginTime, endTime, resp)
	return resp, err
}

// ServerTracing applies tracing to a gRPC server handler.
//
// It takes a context.Context object, a request object of any type,
// a *grpc.UnaryServerInfo object, and a grpc.UnaryHandler function.
//
// It returns a response object of any type and an error.
func ServerTracing(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}
	// 将请求添加到 metadata
	mdmap := metatext.MetadataTextMap{MD: md}
	attrs := []attribute.KeyValue{}

	mdmap.Set("method", info.FullMethod)
	mdmap.Set("remote-addr", md.Get("x-forwarded-for")[0])

	_ = mdmap.ForeachKey(func(key, val string) error {
		attrs = append(attrs, attribute.String(key, val))
		return nil
	})

	tr := global.Tracer.Tracer("grpc")
	spanName := info.FullMethod
	attrs = append(attrs, attribute.String("service", "flag-test-ok"))

	spanOpts := []trace.SpanStartOption{
		trace.WithAttributes(
			attrs...,
		),
	}

	c, span := tr.Start(ctx, spanName, spanOpts...)
	defer span.End()

	return handler(c, req)
}
