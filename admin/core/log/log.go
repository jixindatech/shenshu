package log

import (
	"admin/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Logger *zap.Logger

func getLevel(level string) zapcore.Level {
	res := zapcore.DebugLevel
	switch {
	case level == "debug":
		break
	case level == "info":
		res = zapcore.InfoLevel
	case level == "warn":
		res = zapcore.WarnLevel
	case level == "error":
		res = zapcore.ErrorLevel
	case level == "panic":
		res = zapcore.PanicLevel
	case level == "fatal":
		res = zapcore.FatalLevel
	}

	return res
}

func Setup(log *config.Log) error {
	hook := lumberjack.Logger{
		Filename:   log.Filename,
		MaxSize:    log.MaxSize,
		MaxBackups: log.MaxBackup,
		MaxAge:     log.MaxAge,
		Compress:   log.Compress,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "log",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	level := getLevel(log.Level)
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atomicLevel,
	)

	if level == zapcore.DebugLevel {
		caller := zap.AddCaller()
		development := zap.Development()
		Logger = zap.New(core, caller, development)
	} else {
		Logger = zap.New(core)
	}

	return nil
}
