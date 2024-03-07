package logs

var defaultLogger Logger

func InitDefaultLogger(opts ...Option) error {
	opts = append(opts, WithCallerSkip(2))

	logger, err := NewZapLogger(opts...)
	if err != nil {
		return err
	}

	defaultLogger = logger

	return nil
}
