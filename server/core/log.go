package core

import (
	"ccsu/global"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

func init() {
	var logger *zap.Logger
	core := zapcore.NewCore(newZapEncoder(), newWriteSyncer(), zap.NewAtomicLevelAt(getLevel(global.Config.Log.Level)))

	// set initialize fields
	field := zap.Fields(zap.String("ApplicationName", global.Config.System.Name))

	// judge current environment is pro?
	if global.Config.System.Mode == "dev" {
		// enable dev mode
		caller := zap.AddCaller()
		development := zap.Development()
		logger = zap.New(core, caller, development, field)
	} else {
		logger = zap.New(core, field)
	}

	global.Log = logger
}

func newWriteSyncer() zapcore.WriteSyncer {
	if global.Config.System.Mode == "dev" {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	} else {
		logConfig := global.Config.Log
		now := time.Now()
		hook := lumberjack.Logger{
			Filename:   logConfig.Path + string(os.PathSeparator) + fmt.Sprintf("%04d%02d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()) + ".log",
			MaxSize:    logConfig.MaxSize,
			MaxBackups: logConfig.MaxBackups,
			MaxAge:     logConfig.MaxAge,
			Compress:   logConfig.Compress,
		}
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
	}
}

func newZapEncoder() zapcore.Encoder {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "lineNum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 大写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,  // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	if global.Config.Log.Encoder == "json" {
		return zapcore.NewJSONEncoder(encoderCfg)
	} else {
		return zapcore.NewConsoleEncoder(encoderCfg)
	}
}

func getLevel(l string) zapcore.Level {
	switch strings.ToUpper(l) {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "WARN":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	case "DPANIC":
		return zapcore.DPanicLevel
	case "PANIC":
		return zapcore.PanicLevel
	case "FATAL":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
