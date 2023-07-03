package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/internal/model"
	"github.com/lc-1010/OneBlogService/internal/routers"
	"github.com/lc-1010/OneBlogService/pkg/logger"
	"github.com/lc-1010/OneBlogService/pkg/setting"
	"github.com/lc-1010/OneBlogService/pkg/tracer"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	gin.SetMode(global.ServerSetting.RunMode)

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,  //10 * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeout, //10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("s.ListenAndServe err:%v", err)
	}
}

func init() {

	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err:%v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err:%v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err:%v", err)
	}

}

// setupSetting
func setupSetting() error {
	setting, err := setting.NewSetting("")
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email", &global.EmailSettings)
	if err != nil {
		return err
	}
	//jwt expire 7200 second
	global.JWTSetting.Expire *= time.Second

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	// use = 设置启动，不能用:= 因为其他包调用时是会是nil

	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogServePath + "/" +
		global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupTracer() error {
	tracerProvider, err := tracer.NewJaegerTrancer(
		"blog",
		"127.0.0.1",
		"6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = tracerProvider
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// defer func(ctx context.Context) {
	// 	// Do not make the application hang when it is shutdown.
	// 	ctx, cancel = context.WithTimeout(ctx, time.Second*5)
	// 	defer cancel()
	// 	if err := tracerProvider.Shutdown(ctx); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }(ctx)

	tr := global.Tracer.Tracer("component-main")
	_, span := tr.Start(ctx, "init")
	defer span.End()
	return nil
}
