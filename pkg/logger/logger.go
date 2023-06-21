package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

type Level int8

type Fields map[string]any

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

// 日志标准化 ，具体方法来对日志实例初始化和标准化参数进行绑定

// clone 克隆
func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

// 设置公共字段
func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}
	for k, v := range f {
		ll.fields[k] = v
	}
	return ll
}

// WithContext 设置日志上下文属性
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

// WithCaller 设置某一层调用栈信息
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s:%d %s", file, line, f.Name())}
	}
	return ll
}

// WithCallersFrames 当前调用栈信息
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}

// 日志格式化输出
func (l *Logger) WithTrace() *Logger {
	ginCtx, ok := l.ctx.(*gin.Context)
	if ok {
		return l.WithFields(Fields{
			"trace_id": ginCtx.MustGet("X-Trace-ID"),
			"span_id":  ginCtx.MustGet("X-Span-ID"),
		})
	}
	return l
}

func (l *Logger) JSONFormat(level Level, message string) map[string]any {
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}

func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(body)
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

func (l *Logger) log(ctx context.Context, level Level, v ...any) {
	l.WithContext(ctx).WithTrace().Output(level, fmt.Sprint(v...))
}
func (l *Logger) logf(ctx context.Context, level Level, format string, v ...any) {
	l.WithContext(ctx).WithTrace().Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(ctx context.Context, v ...any) {
	l.log(ctx, LevelDebug, v...)
}
func (l *Logger) Debugf(ctx context.Context, format string, v ...any) {
	l.logf(ctx, LevelDebug, format, v...)
}

func (l *Logger) Info(ctx context.Context, v ...any) {
	l.log(ctx, LevelInfo, v...)
}
func (l *Logger) Infof(ctx context.Context, format string, v ...any) {
	l.logf(ctx, LevelInfo, format, v...)
}

func (l *Logger) Warn(ctx context.Context, v ...any) {
	l.log(ctx, LevelWarn, v...)
}
func (l *Logger) Warnf(ctx context.Context, format string, v ...any) {
	l.logf(ctx, LevelWarn, format, v...)
}

func (l *Logger) Error(ctx context.Context, v ...any) {
	l.log(ctx, LevelError, v...)
}
func (l *Logger) Errorf(ctx context.Context, format string, v ...any) {
	l.logf(ctx, LevelError, format, v...)
}

func (l *Logger) Fata(ctx context.Context, v ...any) {
	l.log(ctx, LevelFatal, v...)
}
func (l *Logger) Fatalf(ctx context.Context, format string, v ...any) {
	l.logf(ctx, LevelFatal, format, v...)
}

func (l *Logger) Panic(ctx context.Context, v ...any) {
	l.log(ctx, LevelPanic, v...)
}
func (l *Logger) Panicf(ctx context.Context, format string, v ...any) {
	l.logf(ctx, LevelPanic, format, v...)
}
