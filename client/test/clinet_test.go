package test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/pkg/tracer"
)

func TestTacer(t *testing.T) {
	TSpan()
}

type b string

func TSpan() {

	tracerPorvider, err := tracer.NewJaegerTrancer(
		"cligrpc",
		"127.0.0.1",
		"6831",
	)
	if err != nil {
		log.Fatal(err)
	}
	global.Tracer = tracerPorvider

	ctx := context.Background()
	tr := global.Tracer.Tracer("test11")
	pp := context.WithValue(ctx, b("ba"), b("j"))
	_, span := tr.Start(pp, "cli-rpc")
	defer span.End()
	fmt.Println("ok tracer")
}
