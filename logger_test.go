package zapx_test

import (
	"context"
	"testing"

	"github.com/xylonx/zapx"
	"go.uber.org/zap"
)

type traceDecoder struct{}

func (decoder traceDecoder) DecodeCtx(ctx context.Context) []zap.Field {
	fields := make([]zap.Field, 0)
	if traceId, ok := ctx.Value("TraceId").(string); ok {
		fields = append(fields, zap.String("TraceId", traceId))
	}
	return fields
}

func TestLog(t *testing.T) {
	l := zap.NewExample()
	logger := zapx.WrapZapLogger(l, nil)

	ctx := context.WithValue(context.Background(), "A", "ababab")
	logger.WithContext(ctx).Info("xxxx")

	ctxLogger := zapx.WrapZapLogger(l, traceDecoder{})
	ctxLogger.WithContext(ctx).Info("xxxx")

	zapx.Info("Hello world!", zap.String("name", "xylonx"))
	zapx.Warn("this is a warn info", zap.String("name", "xylon"))
}
