package logger

// LoggerLevel ...
type Level int8

// LoggerEnv ...
type Env int8

// LeveledLogger defined ther logger formats message according to format specifier
// and writes to log with levels
type LeveledLogger interface {
	Debugf(format string, params ...interface{})

	Infof(format string, params ...interface{})

	Warnf(format string, params ...interface{})

	Errorf(format string, params ...interface{})

	Fatalf(format string, params ...interface{})

	Debug(v ...interface{})

	Info(v ...interface{})

	Warn(v ...interface{})

	Error(v ...interface{})

	Fatal(v ...interface{})
}

type LoggerConfig struct {
	Env   Env
	Name  string
	Level Level
	Path  string
}

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)
const (
	Development Env = 1 + iota
	Stage
	Production
)
