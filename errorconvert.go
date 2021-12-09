package zapx

import (
	"errors"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type errorConverter struct {
	logger *zap.Logger
	buffer *strings.Builder
}

func newErrorConverter() *errorConverter {
	errorConverter := new(errorConverter)
	errorConverter.buffer = &strings.Builder{}

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format("2006-01-02 15:04:05.000000"))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(errorConverter.buffer)),
		zap.NewAtomicLevelAt(zap.ErrorLevel),
	)

	errorConverter.logger = zap.New(core)

	return errorConverter
}

func (ec *errorConverter) clone() *errorConverter {
	logger := *ec.logger.With()
	return &errorConverter{
		logger: &logger,
		buffer: &strings.Builder{},
	}
}

func (ec *errorConverter) fieldsError(msg string, fields ...zap.Field) error {
	ec.logger.Error(msg, fields...)
	ec.logger.Sync() // nolint:errcheck
	err := errors.New(ec.buffer.String())
	ec.buffer.Reset()
	return err
}

func (ec *errorConverter) sugarError(template string, args ...interface{}) error {
	ec.logger.Sugar().Errorf(template, args...)
	ec.logger.Sync() // nolint:errcheck
	err := errors.New(ec.buffer.String())
	ec.buffer.Reset()
	return err
}

func (ec *errorConverter) fieldsDPanic(msg string, fields ...zap.Field) error {
	ec.logger.DPanic(msg, fields...)
	ec.logger.Sync() // nolint:errcheck
	err := errors.New(ec.buffer.String())
	ec.buffer.Reset()
	return err
}

func (ec *errorConverter) sugarDPanic(template string, args ...interface{}) error {
	ec.logger.Sugar().DPanicf(template, args...)
	ec.logger.Sync() // nolint:errcheck
	err := errors.New(ec.buffer.String())
	ec.buffer.Reset()
	return err
}

func (ec *errorConverter) fieldsPanic(msg string, fields ...zap.Field) error {
	ec.logger.Panic(msg, fields...)
	ec.logger.Sync() // nolint:errcheck
	err := errors.New(ec.buffer.String())
	ec.buffer.Reset()
	return err
}

func (ec *errorConverter) sugarPanic(template string, args ...interface{}) error {
	ec.logger.Sugar().Panicf(template, args...)
	ec.logger.Sync() // nolint:errcheck
	err := errors.New(ec.buffer.String())
	ec.buffer.Reset()
	return err
}

func (ec *errorConverter) fieldsFatal(msg string, fields ...zap.Field) error {
	ec.logger.Fatal(msg, fields...)
	ec.logger.Sync() // nolint:errcheck
	err := errors.New(ec.buffer.String())
	ec.buffer.Reset()
	return err
}

func (ec *errorConverter) sugarFatal(template string, args ...interface{}) error {
	ec.logger.Sugar().Fatalf(template, args...)
	ec.logger.Sync() // nolint:errcheck
	err := errors.New(ec.buffer.String())
	ec.buffer.Reset()
	return err
}
