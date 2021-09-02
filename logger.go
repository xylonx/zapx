package zapx

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	NoopCtxDecoder = &noopCtxDecoder{}
)

type WrapLogger struct {
	logger *zap.Logger
	CtxDecoder
}

type CtxDecoder interface {
	DecodeCtx(context.Context) []zap.Field
}

type noopCtxDecoder struct{}

func (decoder *noopCtxDecoder) DecodeCtx(ctx context.Context) []zap.Field {
	return nil
}

func WrapZapLogger(l *zap.Logger, decoder CtxDecoder) *WrapLogger {
	if decoder == nil {
		decoder = NoopCtxDecoder
	}
	return &WrapLogger{
		logger:     l,
		CtxDecoder: decoder,
	}
}

func (log *WrapLogger) clone() *WrapLogger {
	l := *log.logger
	return &WrapLogger{
		logger: &l,
	}
}

func (log *WrapLogger) With(fields ...zap.Field) *WrapLogger {
	if len(fields) == 0 {
		return log
	}
	l := log.clone()
	l.logger = l.logger.With(fields...)
	return l
}

func (log *WrapLogger) WithContext(ctx context.Context) *WrapLogger {
	return log.With(log.DecodeCtx(ctx)...)
}

func (log *WrapLogger) Debug(msg string, fields ...zap.Field) {
	log.logger.Debug(msg, fields...)
}

func (log *WrapLogger) Info(msg string, fields ...zap.Field) {
	log.logger.Info(msg, fields...)
}

func (log *WrapLogger) Warn(msg string, fields ...zap.Field) {
	log.logger.Warn(msg, fields...)
}

func (log *WrapLogger) Error(msg string, fields ...zap.Field) {
	log.logger.Error(msg, fields...)
}

func (log *WrapLogger) DPanic(msg string, fields ...zap.Field) {
	log.logger.DPanic(msg, fields...)
}

func (log *WrapLogger) Panic(msg string, fields ...zap.Field) {
	log.logger.Panic(msg, fields...)
}

func (log *WrapLogger) Fatal(msg string, fields ...zap.Field) {
	log.logger.Fatal(msg, fields...)
}

func (log *WrapLogger) Sync() error {
	return log.logger.Sync()
}

func (log *WrapLogger) Core() zapcore.Core {
	return log.logger.Core()
}
