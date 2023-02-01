package log

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

type Field = zapcore.Field

type Logger interface {
	Debug(msg string)
	DebugC(context context.Context, msg string)
	Debugf(format string, args ...interface{})
	DebugfC(context context.Context, msg string)
	DebugW(msg string, keysAndValues ...interface{})
	DebugWC(context context.Context, msg string, keysAndValues ...interface{})
}

var _ Logger = &zapLogger{}

type zapLogger struct {
	zapLogger *zap.Logger
}
type otherLogger struct {
}

func (z *zapLogger) Debug(msg string) {
	z.zapLogger.Debug(msg)
}

func (z *zapLogger) DebugC(context context.Context, msg string) {
	//TODO implement me
	panic("implement me")
}

func (z *zapLogger) Debugf(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (z *zapLogger) DebugfC(context context.Context, msg string) {
	//TODO implement me
	panic("implement me")
}

func (z *zapLogger) DebugW(msg string, keysAndValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (z *zapLogger) DebugWC(context context.Context, msg string, keysAndValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}

var (
	defaultLogger = New(NewOptions())
	mu            sync.Mutex
)

func Debug(msg string) {
	defaultLogger.Debug(msg)
}

func New(opts *Options) *zapLogger {
	if opts == nil {
		opts = NewOptions()
	}

	//实例化zap
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	loggerConfig := zap.Config{
		Level: zap.NewAtomicLevelAt(zapLevel),
	}
	l, err := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel))
	if err != nil {
		panic(err)
	}
	logger := &zapLogger{
		zapLogger: l.Named(opts.Name),
	}
	return logger
}

func Init(opt *Options) {
	//看起来没有问题， 并发问题, 因为我们后面可能希望我们的这个全局的logger是动态的
	mu.Lock()
	defer mu.Unlock()
	defaultLogger = New(opt)
}
