package zapx

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger         *WrapLogger
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
	logger = &WrapLogger{
		logger:     l,
		CtxDecoder: decoder,
	}
	return logger
}

func Use(l *zap.Logger, decoder CtxDecoder) {
	logger = &WrapLogger{
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

func With(fields ...zap.Field) *WrapLogger {
	if len(fields) == 0 {
		return logger
	}
	l := logger.clone()
	l.logger = l.logger.With(fields...)
	return l
}

func (log *WrapLogger) WithContext(ctx context.Context) *WrapLogger {
	return log.With(log.DecodeCtx(ctx)...)
}

func WithContext(ctx context.Context) *WrapLogger {
	return logger.With(logger.DecodeCtx(ctx)...)
}

func (log *WrapLogger) Debug(msg string, fields ...zap.Field) {
	log.logger.Debug(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func (log *WrapLogger) Info(msg string, fields ...zap.Field) {
	log.logger.Info(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func (log *WrapLogger) Warn(msg string, fields ...zap.Field) {
	log.logger.Warn(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func (log *WrapLogger) Error(msg string, fields ...zap.Field) {
	log.logger.Error(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func (log *WrapLogger) DPanic(msg string, fields ...zap.Field) {
	log.logger.DPanic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}

func (log *WrapLogger) Panic(msg string, fields ...zap.Field) {
	log.logger.Panic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

func (log *WrapLogger) Fatal(msg string, fields ...zap.Field) {
	log.logger.Fatal(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func (log *WrapLogger) Sync() error {
	return log.logger.Sync()
}

func Sync() error {
	return logger.Sync()
}

func (log *WrapLogger) Core() zapcore.Core {
	return log.logger.Core()
}

func Core() zapcore.Core {
	return logger.Core()
}
