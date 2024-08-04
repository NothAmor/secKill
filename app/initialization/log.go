package initialization

import (
	"os"
	"secKill/app/common"

	"github.com/charmbracelet/log"
)

func initLog() {
	logConfig := common.Config.Log
	level, err := log.ParseLevel(logConfig.Level)
	if err != nil {
		level = log.DebugLevel
		log.Printf("log level parse error: [%v], use default level: [debug]", err)
	}

	common.Logger = log.NewWithOptions(os.Stderr, log.Options{
		Level:           level,                    // 日志级别
		Prefix:          logConfig.Prefix,         // 日志前缀
		ReportTimestamp: logConfig.Timestamp,      // 是否显示时间戳
		ReportCaller:    logConfig.Caller,         // 是否显示调用信息
		TimeFunction:    log.NowUTC,               // 时间函数
		Formatter:       log.JSONFormatter,        // 日志格式
		CallerFormatter: log.ShortCallerFormatter, // 调用信息格式
	})
}
