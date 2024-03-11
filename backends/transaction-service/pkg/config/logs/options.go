package logs

import (
	"encoding/json"
	"stori/transaction-service/pkg/config/errors"
	"time"

	"go.uber.org/zap"
)

type Option func(*options)

type Field = zap.Field

type options struct {
	level      Level
	encoding   string
	callerSkip int
	fields     []Field
}

func (o *options) apply(opts ...Option) *options {
	cloned := o.clone()

	for _, opt := range opts {
		opt(cloned)
	}

	return cloned
}

func (o *options) clone() *options {
	cloned := *o
	return &cloned
}

func WithLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}

func WithEncoding(encoding string) Option {
	return func(o *options) {
		o.encoding = encoding
	}
}

func WithCallerSkip(skip int) Option {
	return func(o *options) {
		o.callerSkip = skip
	}
}

func WithBool(key string, value bool) Option {
	return func(o *options) {
		o.fields = append(o.fields, zap.Bool(key, value))
	}
}

func WithBinary(key string, value []byte) Option {
	return func(o *options) {
		o.fields = append(o.fields, zap.Binary(key, value))
	}
}

func WithDuration(key string, value time.Duration) Option {
	return func(o *options) {
		o.fields = append(o.fields, zap.Duration(key, value))
	}
}

func WithError(err error) Option {
	return func(o *options) {
		if err, ok := err.(*errors.Error); ok {
			b, err := err.MarshalJSON()
			if err == nil {
				o.fields = append(o.fields, zap.Any("error", json.RawMessage(b)))
				return
			}
		}

		o.fields = append(o.fields, zap.Error(err))
	}
}

func WithFloat32(key string, value float32) Option {
	return func(o *options) {
		o.fields = append(o.fields, zap.Float32(key, value))
	}
}

func WithFloat64(key string, value float64) Option {
	return func(o *options) {
		o.fields = append(o.fields, zap.Float64(key, value))
	}
}

func WithInt(key string, value int) Option {
	return func(o *options) {
		o.fields = append(o.fields, zap.Int(key, value))
	}
}

func WithInt32(key string, value int32) Option {
	return func(o *options) {
		o.fields = append(o.fields, zap.Int32(key, value))
	}
}

func WithInt64(key string, value int64) Option {
	return func(o *options) {
		o.fields = append(o.fields, zap.Int64(key, value))
	}
}

func WithStack(key string) Option {
	return func(o *options) {
		o.fields = append(o.fields, zap.Stack(key))
	}
}

func WithString(key string, value string) Option {
	return func(o *options) {
		o.fields = append(o.fields, zap.String(key, value))
	}
}

func WithTime(key string, value time.Time) Option {
	return func(o *options) {
		o.fields = append(o.fields, zap.Time(key, value))
	}
}

func WithObject(key string, value interface{}) Option {
	return func(o *options) {
		o.fields = append(o.fields, zap.Any(key, value))
	}
}
