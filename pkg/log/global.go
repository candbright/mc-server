package log

import (
	"go.uber.org/zap"
)

var globalConfig *Config
var loggers = make(map[string]*zap.Logger)

func Set(cfg *Config) {
	globalConfig = cfg
}

func Ins() *zap.Logger {
	if globalConfig == nil {
		panic("please set the global config by using Set() method")
	}
	if loggers[globalConfig.Path] != nil {
		return loggers[globalConfig.Path]
	} else {
		return New(&Config{
			Path:       globalConfig.Path,
			Level:      globalConfig.Level,
			TraceLevel: globalConfig.TraceLevel,
		})
	}
}

func Debug(format string, v ...interface{}) {
	Ins().Sugar().Debugf(format, v...)
}

func Info(format string, v ...interface{}) {
	Ins().Sugar().Infof(format, v...)
}

func Warn(format string, v ...interface{}) {
	Ins().Sugar().Warnf(format, v...)
}

func Error(format string, v ...interface{}) {
	Ins().Sugar().Errorf(format, v...)
}
