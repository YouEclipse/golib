package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var _ LeveledLogger = &ZapLogger{}

type ZapLogger struct {
	logger *zap.SugaredLogger
	level  *zapcore.Level
}

func NewZapLogger(cfg *LoggerConfig) *ZapLogger {
	wrap := &ZapLogger{}
	zapLevel := zapcore.Level(cfg.Level)

	var (
		encoderConf zapcore.EncoderConfig
		cores       = make([]zapcore.Core, 0)
	)

	encoderConf = zap.NewProductionEncoderConfig()
	encoderConf.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cores = append(cores, zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConf),
		zapcore.AddSync(os.Stdout),
		zapcore.Level(zapLevel),
	))

	if cfg.Env == Development {
		//in development environment, only logging to stdout
		//so do nothing here
	} else if cfg.Env == Stage {
		//in stage environment, logging to stdout and files
		fLogger := &lumberjack.Logger{
			Filename:   cfg.Path + cfg.Name + ".log",
			MaxSize:    10,
			MaxBackups: 10,
			MaxAge:     30,
			Compress:   false,
		}
		cores = append(cores, zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConf),
			zapcore.AddSync(fLogger),
			zapcore.Level(zapLevel),
		))
	} else {
		//in production environment, logging to stdout and efk-cluster
		cores = append(cores, zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConf),
			zapcore.AddSync(&EFKLoggerSyncer{}),
			zapcore.Level(zapLevel),
		))
	}

	core := zapcore.NewTee(cores...)

	wrap.logger = zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(2),
		//zap.AddStacktrace(zapLevel),
	).With().Sugar()

	return wrap
}
func (l *ZapLogger) Debugf(format string, params ...interface{}) {
	defer l.logger.Sync()
	l.logger.Debugf(format, params...)
}

func (l *ZapLogger) Infof(format string, params ...interface{}) {
	l.logger.Infof(format, params...)
}

func (l *ZapLogger) Warnf(format string, params ...interface{}) {
	l.logger.Warnf(format, params...)
}

func (l *ZapLogger) Errorf(format string, params ...interface{}) {
	l.logger.Errorf(format, params...)
}

func (l *ZapLogger) Fatalf(format string, params ...interface{}) {
	l.logger.Fatalf(format, params...)
}

func (l *ZapLogger) Debug(v ...interface{}) {
	l.logger.Debug(v...)
}

func (l *ZapLogger) Info(v ...interface{}) {
	l.logger.Info(v...)
}

func (l *ZapLogger) Warn(v ...interface{}) {
	l.logger.Warn(v...)
}

func (l *ZapLogger) Error(v ...interface{}) {
	l.logger.Error(v...)
}

func (l *ZapLogger) Fatal(v ...interface{}) {
	l.logger.Fatal(v...)
}
