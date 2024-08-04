package logger

import (
	"context"
	"os"
	"secKill/app/common"
	"secKill/app/config"
	"testing"

	"github.com/charmbracelet/log"
)

func testInit() {
	common.Config = &config.Config{
		Log: config.Log{
			Level:     "debug",
			Caller:    true,
			Timestamp: true,
			Prefix:    "test",
		},
	}

	logConfig := common.Config.Log
	common.Logger = log.NewWithOptions(os.Stderr, log.Options{
		Level:           log.DebugLevel,           // 日志级别
		Prefix:          logConfig.Prefix,         // 日志前缀
		ReportTimestamp: logConfig.Timestamp,      // 是否显示时间戳
		ReportCaller:    logConfig.Caller,         // 是否显示调用信息
		TimeFunction:    log.NowUTC,               // 时间函数
		Formatter:       log.JSONFormatter,        // 日志格式
		CallerFormatter: log.ShortCallerFormatter, // 调用信息格式
	})
}

func TestIx(t *testing.T) {
	tag := "testIx"
	testInit()

	c := context.Background()

	// empty context
	Ix(c, tag, "test msg: [%+v]", "test")

	// context with values
	// tradeNode := GenLogTraceMetadata()
	// ctx := context.WithValue(c, GetMetadataKey(), tradeNode)
	// Ix(ctx, tag, "test context with values: [%+v]", "test")
}
