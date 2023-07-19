package blog_api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/pkg/metatext"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
)

const (
	APP_KEY    = "APP_KEY"
	APP_SECRET = "APP_SECRET"
)

// API blog api sdk
type API struct {
	URL    string
	Client *http.Client
}

// AccessToken
type AccessToken struct {
	Token string `json:"token"`
}

// NewAPI creates a new API instance.
//
// It takes a URL string as a parameter and returns a pointer to an API object.
// 实现http 请求到具体路由
func NewAPI(url string) *API {
	return &API{URL: url, Client: http.DefaultClient}
}

func (a *API) getAccessToken(c context.Context) (string, error) {

	body, err := a.httpGet(c, fmt.Sprintf("%s?app_key=%s&app_secret=%s", "auth", APP_KEY, APP_SECRET))
	if err != nil {
		return "", err
	}
	var accessToken AccessToken
	_ = json.Unmarshal(body, &accessToken)
	return accessToken.Token, nil
}

// httpGet get function
func (a *API) httpGet(ctx context.Context, path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", a.URL, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	// 将请求添加到 metadata
	mdmap := metatext.MetadataTextMap{MD: md}
	attrs := []attribute.KeyValue{}

	//mdmap.Set("method", info.FullMethod)
	//mdmap.Set("remote-addr", md.Get("x-forwarded-for")[0])

	_ = mdmap.ForeachKey(func(key, val string) error {
		attrs = append(attrs, attribute.String(key, val))
		return nil
	})

	attrs = append(attrs, attribute.String("api", "md"),
		attribute.String("path", path),
		attribute.String("method", req.Method),
	)

	tracer := global.Tracer.Tracer("api")
	ctx, span := tracer.Start(ctx, url, trace.WithAttributes(
	//attrs...,
	//	attribute.String("client-id", md.Get("client_id")[0]),
	))

	propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	propagator.Inject(ctx, propagation.HeaderCarrier(req.Header))

	client := http.Client{Timeout: 5 * time.Second}
	startTime := time.Now()
	resp, err := client.Do(req)
	endTime := time.Now()

	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	defer resp.Body.Close()
	defer span.End()

	body, _ := io.ReadAll(resp.Body)
	span.SetAttributes(
		attribute.String("status", fmt.Sprintf("%d", resp.StatusCode)),
		attribute.Float64("latency", endTime.Sub(startTime).Seconds()),
	)

	return body, nil
}

func (a *API) GetTagList(ctx context.Context, name string) ([]byte, error) {
	token, err := a.getAccessToken(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	body, err := a.httpGet(ctx, fmt.Sprintf("%s?token=%s&name=%s", "api/v1/tags", token, name))
	if err != nil {
		return nil, err
	}
	return body, nil
}
