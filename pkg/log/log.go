package log

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Path       string
	Level      func() string
	TraceLevel func() string
}

func New(cfg *Config) *zap.Logger {
	if cfg.Level == nil {
		cfg.Level = func() string {
			return "info"
		}
	}
	if cfg.TraceLevel == nil {
		cfg.TraceLevel = func() string {
			return "panic"
		}
	}
	var writeSyncer zapcore.WriteSyncer
	file, err := Attach(cfg.Path, 0755, 0644)
	if err != nil {
		fmt.Printf("%+v\n", err)
		writeSyncer = os.Stdout
	} else {
		writeSyncer = zapcore.NewMultiWriteSyncer(os.Stdout, file)
	}
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	lvlEnablerFactory := func(lvlFunc func() string) zapcore.LevelEnabler {
		return zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			var lvl zapcore.Level
			lvlStr := lvlFunc()
			lvlInt, err := strconv.Atoi(lvlStr)
			if err != nil {
				lvl = zapcore.Level(lvlInt)
				e := (&lvl).UnmarshalText([]byte(lvlStr))
				if e != nil {
					fmt.Printf("%+v\n", e)
				}
			}
			return level >= lvl
		})
	}
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), writeSyncer, lvlEnablerFactory(cfg.Level))
	logger := zap.New(core, zap.AddStacktrace(lvlEnablerFactory(cfg.TraceLevel)))
	return logger
}

func Attach(path string, dirPerm, filePerm os.FileMode) (*os.File, error) {
	if lastIndex := strings.LastIndex(path, "/"); lastIndex != -1 {
		err := os.MkdirAll(path[:lastIndex+1], dirPerm)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, filePerm)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return f, nil
}
