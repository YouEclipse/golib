package logger

import (
	"strings"
	"sync"
)

var defalutLogger LeveledLogger

func InitLogger(cfg *LoggerConfig) {
	if cfg.Name == "" {
		cfg.Name = "logger"
	}
	if cfg.Env <= 0 || cfg.Env > 3 {
		cfg.Env = Development
	}
	if cfg.Path == "" {
		cfg.Path = "/tmp/"
	}

	if !strings.HasSuffix(cfg.Path, "/") {
		cfg.Path += "/"
	}

	once := sync.Once{}
	once.Do(
		func() {
			defalutLogger = NewZapLogger(cfg)
		},
	)
}

func GetLogger() LeveledLogger {
	return defalutLogger
}

func Debugf(format string, params ...interface{}) {
	defalutLogger.Debugf(format, params...)
}

func Infof(format string, params ...interface{}) {
	defalutLogger.Infof(format, params...)
}

func Warnf(format string, params ...interface{}) {
	defalutLogger.Warnf(format, params...)
}

func Errorf(format string, params ...interface{}) {
	defalutLogger.Errorf(format, params...)
}

func Fatalf(format string, params ...interface{}) {
	defalutLogger.Fatalf(format, params...)
}

func Debug(v ...interface{}) {
	defalutLogger.Debug(v...)
}

func Info(v ...interface{}) {
	defalutLogger.Info(v...)
}
func Warn(v ...interface{}) {
	defalutLogger.Warn(v...)
}
func Error(v ...interface{}) {
	defalutLogger.Error(v...)
}
func Fatal(v ...interface{}) {
	defalutLogger.Fatal(v...)
}
