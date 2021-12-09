package zapx

import (
	"context"
	"errors"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/natefinch/lumberjack"
	"github.com/xylonx/zapx/decoder"
)

var (
	// the in-package Logger.
	// It can be overwritten by Use() function
	// By default, it is initialized with zap.NewExample() without any context info.
	_logger *Logger
)

type Logger struct {
	decoder.CtxDecoder

	ctx    context.Context
	fields map[string]string

	opts   *Option
	logger *zap.Logger

	errConverter *errorConverter
}

type Option struct {
	LogSyncer []zapcore.WriteSyncer

	// lumberjack log rolling config.
	// to see more information about it, see its github(https://github.com/natefinch/lumberjack)
	RollingOption *lumberjack.Logger
	LogLevel      string
}

func init() {
	_logger = &Logger{
		CtxDecoder: &decoder.NoopCtxDecoder{},

		ctx:    context.Background(),
		fields: make(map[string]string),

		opts:         &Option{},
		logger:       zap.NewExample(),
		errConverter: newErrorConverter(),
	}
}

func Use(opt *Option) error {
	logger, err := NewLogger(opt)
	if err != nil {
		return err
	}
	_logger = logger
	return nil
}

func NewLogger(opt *Option) (logger *Logger, err error) {
	if opt == nil {
		return nil, errors.New("config can't be nil")
	}

	var zapLogger *zap.Logger

	zapLogger = initZapLogger(opt)

	// with default fields
	zapLogger = zapLogger.With(defaultFields()...)

	logger = &Logger{
		CtxDecoder: &decoder.NoopCtxDecoder{},

		ctx:    context.Background(),
		fields: make(map[string]string),

		opts:   opt,
		logger: zapLogger,

		errConverter: newErrorConverter(),
	}

	return logger, nil
}

// initZapLogger create a zap logger
func initZapLogger(opt *Option) *zap.Logger {
	syncWriters := []zapcore.WriteSyncer{zapcore.AddSync(os.Stdout)}

	if opt.RollingOption != nil {
		syncWriters = append(syncWriters, getRollingLogWriter(opt.RollingOption))
	}

	if opt.LogSyncer != nil {
		syncWriters = append(syncWriters, opt.LogSyncer...)
	}

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format("2006-01-02 15:04:05.000000"))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		zapcore.NewMultiWriteSyncer(syncWriters...),
		zap.NewAtomicLevelAt(getLoggerLevel(opt.LogLevel)),
	)

	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// getRollingLogWriter generate zap.WriteSyncer to store log into file with auto rotate
func getRollingLogWriter(opt *lumberjack.Logger) zapcore.WriteSyncer {
	if opt.Filename == "" {
		opt.Filename = "./log/zapx.log"
	}
	if opt.MaxSize == 0 {
		opt.MaxSize = 100
	}
	if opt.MaxAge == 0 {
		opt.MaxAge = 30
	}

	return zapcore.AddSync(opt)
}

func (log *Logger) clone() *Logger {
	l := *log.logger

	f := make(map[string]string)
	for k := range log.fields {
		f[k] = log.fields[k]
	}

	return &Logger{
		CtxDecoder: log.CtxDecoder,

		ctx:    log.ctx,
		fields: f,

		opts:   log.opts,
		logger: &l,

		errConverter: log.errConverter.clone(),
	}
}

func (log *Logger) WithContext(ctx context.Context) *Logger {
	l := log.With(log.DecodeCtx(ctx)...)
	l.ctx = ctx
	return l
}

func WithContext(ctx context.Context) *Logger {
	return _logger.WithContext(ctx)
}

// With creates a child logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
func (log *Logger) With(fields ...zap.Field) *Logger {
	if len(fields) == 0 {
		return log
	}
	l := log.clone()
	l.logger = l.logger.With(fields...)
	return l
}

func With(fields ...zap.Field) *Logger {
	return _logger.With(fields...)
}

func (log *Logger) Debug(msg string, fields ...zap.Field) {
	log.logger.Debug(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	_logger.Debug(msg, fields...)
}

func (log *Logger) Debugf(template string, args ...interface{}) {
	log.logger.Sugar().Debugf(template, args...)
}

func Debugf(template string, args ...interface{}) {
	_logger.Debugf(template, args...)
}

func (log *Logger) Info(msg string, fields ...zap.Field) {
	log.logger.Info(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	_logger.Info(msg, fields...)
}

func (log *Logger) Infof(template string, args ...interface{}) {
	log.logger.Sugar().Infof(template, args...)
}

func Infof(template string, args ...interface{}) {
	_logger.Infof(template, args...)
}

func (log *Logger) Warn(msg string, fields ...zap.Field) {
	log.logger.Warn(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	_logger.Warn(msg, fields...)
}

func (log *Logger) Warnf(template string, args ...interface{}) {
	log.logger.Sugar().Warnf(template, args...)
}

func Warnf(template string, args ...interface{}) {
	_logger.Warnf(template, args...)
}

func (log *Logger) Error(msg string, fields ...zap.Field) {
	log.logger.Error(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	_logger.Error(msg, fields...)
}

func (log *Logger) Errorf(template string, args ...interface{}) {
	log.logger.Sugar().Errorf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	_logger.Errorf(template, args...)
}

func (log *Logger) DPanic(msg string, fields ...zap.Field) {
	log.logger.DPanic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	_logger.DPanic(msg, fields...)
}

func (log *Logger) DPanicf(template string, args ...interface{}) {
	log.logger.Sugar().DPanicf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	_logger.DPanicf(template, args...)
}

func (log *Logger) Panic(msg string, fields ...zap.Field) {
	log.logger.Panic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	_logger.Panic(msg, fields...)
}

func (log *Logger) Panicf(template string, args ...interface{}) {
	log.logger.Sugar().Panicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	_logger.Panicf(template, args...)
}

func (log *Logger) Fatal(msg string, fields ...zap.Field) {
	log.logger.Fatal(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	_logger.Fatal(msg, fields...)
}

func (log *Logger) Fatalf(template string, args ...interface{}) {
	log.logger.Sugar().Fatalf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	_logger.Fatalf(template, args...)
}
