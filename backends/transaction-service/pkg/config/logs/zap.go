package logs

import (
	config "stori/transaction/pkg/config/common"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ Logger = (*ZapLogger)(nil)

type ZapLogger struct {
	logger  *zap.Logger
	options *options
}

func NewZapLogger(opts ...Option) (*ZapLogger, error) {
	commonCfg := config.Common()

	options := &options{
		level:    InfoLevel,
		encoding: "json",
		fields: []Field{
			zap.String("env", commonCfg.Env.String()),
			zap.String("service", commonCfg.Service),
		},
		callerSkip: 1,
	}
	options = options.apply(opts...)

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.Level(options.level)),
		Development: commonCfg.Env == "dev",
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         options.encoding,
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &ZapLogger{
		logger:  logger.WithOptions(zap.AddCallerSkip(options.callerSkip)),
		options: options,
	}, nil
}

func (l *ZapLogger) Info(msg string, opts ...Option) {
	options := l.options.apply(opts...)

	if options.level > InfoLevel {
		return
	}

	l.logger.Info(msg, options.fields...)
}

func (l *ZapLogger) Warn(msg string, opts ...Option) {
	options := l.options.apply(opts...)

	if options.level > WarnLevel {
		return
	}

	l.logger.Warn(msg, options.fields...)
}

func (l *ZapLogger) Error(msg string, opts ...Option) {
	options := l.options.apply(opts...)

	if options.level > ErrorLevel {
		return
	}

	l.logger.Error(msg, options.fields...)
}

func (l *ZapLogger) Debug(msg string, opts ...Option) {
	options := l.options.apply(opts...)

	if options.level > DebugLevel {
		return
	}

	l.logger.Debug(msg, options.fields...)
}

func (l *ZapLogger) Fatal(msg string, opts ...Option) {
	options := l.options.apply(opts...)

	if options.level > FatalLevel {
		return
	}

	l.logger.Fatal(msg, options.fields...)
}

func (l *ZapLogger) Panic(msg string, opts ...Option) {
	options := l.options.apply(opts...)

	if options.level > PanicLevel {
		return
	}

	l.logger.Panic(msg, options.fields...)
}

func (l *ZapLogger) CloneWithOptions(opts ...Option) Logger {
	logger := l.logger
	options := l.options.apply(opts...)

	if callerSkipDiff := options.callerSkip - l.options.callerSkip; callerSkipDiff != 0 {
		logger = logger.WithOptions(zap.AddCallerSkip(callerSkipDiff))
	}

	return &ZapLogger{
		logger:  logger,
		options: options,
	}
}
