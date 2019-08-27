package xlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Log struct {
	*zap.Logger
}

type Config struct {
	File  string
	Level string
}

func New(c *Config) *Log {
	var w zapcore.WriteSyncer
	if c.File != "" {
		w = zapcore.AddSync(&lumberjack.Logger{
			Filename: c.File,
			MaxSize:  128,
			//MaxBackups: 10,
			//MaxAge:     7,
			LocalTime: true,
			Compress:  true,
		})
	} else {
		w = os.Stdout
	}

	e := zap.NewProductionEncoderConfig()
	e.EncodeTime = zapcore.ISO8601TimeEncoder

	var level zapcore.Level
	switch c.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(e),
		w,
		level,
	)

	return &Log{zap.New(core, zap.AddCaller())}
}

var DefaultLog *Log
var DefaultSugaredLogger *zap.SugaredLogger

func init() {
	DefaultLog = &Log{New(&Config{"", "debug"}).WithOptions(zap.AddCallerSkip(1))}
	DefaultSugaredLogger = DefaultLog.Sugar()
}

func Init(c *Config) *Log {
	DefaultLog = &Log{New(c).WithOptions(zap.AddCallerSkip(1))}
	DefaultSugaredLogger = DefaultLog.Sugar()
	return DefaultLog
}

func Debug(args ...interface{}) {
	DefaultSugaredLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	DefaultSugaredLogger.Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	DefaultSugaredLogger.Debugw(msg, keysAndValues...)
}

func Debugz(msg string, fields ...zap.Field) {
	DefaultLog.Debug(msg, fields...)
}

func Info(args ...interface{}) {
	DefaultSugaredLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	DefaultSugaredLogger.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	DefaultSugaredLogger.Infow(msg, keysAndValues...)
}

func Infoz(msg string, fields ...zap.Field) {
	DefaultLog.Info(msg, fields...)
}

func Warn(args ...interface{}) {
	DefaultSugaredLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	DefaultSugaredLogger.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	DefaultSugaredLogger.Warnw(msg, keysAndValues...)
}

func Warnz(msg string, fields ...zap.Field) {
	DefaultLog.Warn(msg, fields...)
}

func Error(args ...interface{}) {
	DefaultSugaredLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	DefaultSugaredLogger.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	DefaultSugaredLogger.Errorw(msg, keysAndValues...)
}

func Errorz(msg string, fields ...zap.Field) {
	DefaultLog.Error(msg, fields...)
}

func Panic(args ...interface{}) {
	DefaultSugaredLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	DefaultSugaredLogger.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	DefaultSugaredLogger.Panicw(msg, keysAndValues...)
}

func Panicz(msg string, fields ...zap.Field) {
	DefaultLog.Panic(msg, fields...)
}

func Fatal(args ...interface{}) {
	DefaultSugaredLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	DefaultSugaredLogger.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	DefaultSugaredLogger.Fatalw(msg, keysAndValues...)
}

func Fatalz(msg string, fields ...zap.Field) {
	DefaultLog.Fatal(msg, fields...)
}

func WithOptions(opts ...zap.Option) *Log {
	return &Log{DefaultLog.Logger.WithOptions(opts...)}
}
