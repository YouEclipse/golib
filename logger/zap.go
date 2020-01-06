package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var _ LeveledLogger = &ZapLogger{}

type ZapLogger struct {
	logger *zap.SugaredLogger
	level  *zapcore.Level
}

func NewZapLogger(level LoggerLevel, env LoggerEnv) *ZapLogger {
	_logger := &ZapLogger{}
	zapLevel := zapcore.Level(level)

	var (
		encoderConf zapcore.EncoderConfig
		encoder     zapcore.Encoder
		syncer      zapcore.WriteSyncer
	)
	if env == Development {
		//in development environment, logging in stdout
		encoderConf = zap.NewDevelopmentEncoderConfig()
		encoderConf.EncodeTime = zapcore.RFC3339TimeEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConf)
		syncer = zapcore.AddSync(os.Stdout)
	} else {
		//in stage environment, logging in stdout and files
		encoderConf = zap.NewProductionEncoderConfig()
		encoderConf.EncodeTime = zapcore.RFC3339TimeEncoder
		encoder = zapcore.NewJSONEncoder(encoderConf)
		fLogger := &lumberjack.Logger{
			Filename:   "./test.log",
			MaxSize:    10,
			MaxBackups: 5,
			MaxAge:     30,
			Compress:   false,
		}
		syncer = zapcore.AddSync(fLogger)
	}

	core := zapcore.NewCore(
		encoder,
		syncer,
		zapcore.Level(zapLevel),
	)

	_logger.logger = zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(2),
		zap.AddStacktrace(zapLevel),
	).With().Sugar()

	return _logger
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
