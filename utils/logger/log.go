package logger

import (
	"context"
	"secKill/app/common"

	"github.com/charmbracelet/log"
)

func Ix(ctx context.Context, tag, msg string, fields ...interface{}) {
	build(ctx, tag, log.InfoLevel, msg, fields...)
}

func I(tag, msg string, fields ...interface{}) {
	build(context.Background(), tag, log.InfoLevel, msg, fields...)
}

func Wx(ctx context.Context, tag, msg string, fields ...interface{}) {
	build(ctx, tag, log.WarnLevel, msg, fields...)
}

func W(tag, msg string, fields ...interface{}) {
	build(context.Background(), tag, log.WarnLevel, msg, fields...)
}

func Ex(ctx context.Context, tag, msg string, fields ...interface{}) {
	build(ctx, tag, log.ErrorLevel, msg, fields...)
}

func E(tag, msg string, fields ...interface{}) {
	build(context.Background(), tag, log.ErrorLevel, msg, fields...)
}

func Dx(ctx context.Context, tag, msg string, fields ...interface{}) {
	build(ctx, tag, log.DebugLevel, msg, fields...)
}

func D(tag, msg string, fields ...interface{}) {
	build(context.Background(), tag, log.DebugLevel, msg, fields...)
}

func Fx(ctx context.Context, tag, msg string, fields ...interface{}) {
	build(ctx, tag, log.FatalLevel, msg, fields...)
}

func F(tag, msg string, fields ...interface{}) {
	build(context.Background(), tag, log.FatalLevel, msg, fields...)
}

func build(ctx context.Context, tag string, level log.Level, msg string, fields ...interface{}) {
	var loggerFields []interface{}

	// 添加tag信息
	loggerFields = append(loggerFields, "tag", tag)

	// 链路追踪信息
	traceMetaData := ExtractTraceNodeFromContext(ctx)
	if traceMetaData != nil {
		loggerFields = append(loggerFields, GetTraceIdKey(), traceMetaData.Get(GetTraceIdKey()))
	}

	switch level {
	case log.InfoLevel:
		common.Logger.With(loggerFields...).Infof(msg, fields...)
	case log.DebugLevel:
		common.Logger.With(loggerFields...).Debugf(msg, fields...)
	case log.WarnLevel:
		common.Logger.With(loggerFields...).Warnf(msg, fields...)
	case log.ErrorLevel:
		common.Logger.With(loggerFields...).Errorf(msg, fields...)
	case log.FatalLevel:
		common.Logger.With(loggerFields...).Fatalf(msg, fields...)
	}
}
