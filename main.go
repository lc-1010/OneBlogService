package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"
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

var (
	port    string
	runMode string
	config  string

	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitID  string
)

// go build -ldflags "-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.gitCommitID=`git rev-parse --short HEAD` -X main.buildVersion=1.0.0"
func main() {
	if isVersion {
		fmt.Printf("build_time: %s\n", buildTime)
		fmt.Printf("build_version: %s\n", buildVersion)
		fmt.Printf("git_commit_id: %s\n", gitCommitID)
	}
	gin.SetMode(global.ServerSetting.RunMode)

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HTTPPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,  //10 * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeout, //10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err:%v", err)
		}
	}()

	// 使用go channel 实现中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Tracer 推出
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		if err := global.Tracer.Shutdown(ctx); err != nil {
			log.Fatal(err)
		} else {
			log.Println("Tracer shutdown")
		}
	}(ctx)
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")

}

func init() {
	// 处理 flag provided but not defined: -test.paniconexit0
	testing.Init() //或者把 flag.Parase放到 māin_中
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err:%v", err)
	}
	err = setupSetting()
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
	setting, err := setting.NewSetting(strings.Split(config, ",")...)
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

	// set run mode and port
	if port != "" {
		global.ServerSetting.HTTPPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
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
		"myblog",
		"127.0.0.1",
		"6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = tracerProvider
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tr := global.Tracer.Tracer("component-main")
	_, span := tr.Start(ctx, "init")
	defer span.End()
	return nil
}

func setupFlag() error {

	flag.BoolVar(&isVersion, "version", false, "show version")

	flag.StringVar(&port, "port", "8000", "启动端口")
	flag.StringVar(&runMode, "mode", "debug", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定配置文件路径")
	if !flag.Parsed() {
		flag.Parse()
	}
	return nil
}
