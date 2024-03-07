package logs

import (
	"context"
)

type loggerKey struct{}

func getLogger(ctx context.Context) Logger {
	if logger := fromContext(ctx); logger != nil {
		return logger
	}

	if defaultLogger != nil {
		return defaultLogger
	}

	return nil
}

func fromContext(ctx context.Context) Logger {
	if logger, ok := ctx.Value(loggerKey{}).(Logger); ok {
		return logger
	}

	return nil
}

func ContextWithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger.CloneWithOptions(WithCallerSkip(2)))
}

func AddOptions(ctx context.Context, opts ...Option) context.Context {
	logger := getLogger(ctx)
	if logger == nil {
		return ctx
	}

	return context.WithValue(ctx, loggerKey{}, logger.CloneWithOptions(opts...))
}

func Info(ctx context.Context, messsage string, opts ...Option) {
	logger := getLogger(ctx)
	if logger == nil {
		return
	}

	logger.Info(messsage, opts...)
}

func Warn(ctx context.Context, messsage string, opts ...Option) {
	logger := getLogger(ctx)
	if logger == nil {
		return
	}

	logger.Warn(messsage, opts...)
}

func Error(ctx context.Context, messsage string, opts ...Option) {
	logger := getLogger(ctx)
	if logger == nil {
		return
	}

	logger.Error(messsage, opts...)
}

func Debug(ctx context.Context, messsage string, opts ...Option) {
	logger := getLogger(ctx)
	if logger == nil {
		return
	}

	logger.Debug(messsage, opts...)
}

func Fatal(ctx context.Context, messsage string, opts ...Option) {
	logger := getLogger(ctx)
	if logger == nil {
		return
	}

	logger.Fatal(messsage, opts...)
}

func Panic(ctx context.Context, messsage string, opts ...Option) {
	logger := getLogger(ctx)
	if logger == nil {
		return
	}

	logger.Panic(messsage, opts...)
}
