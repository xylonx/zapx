package zapxdecoder

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var (
	OpentelemetaryDecoder opentelemetaryDecoder
)

type opentelemetaryDecoder struct{}

func (decoder opentelemetaryDecoder) DecodeCtx(ctx context.Context) []zap.Field {
	sc := trace.SpanContextFromContext(ctx)
	return []zap.Field{
		zap.String("TraceID", sc.TraceID().String()),
		zap.String("SpanID", sc.SpanID().String()),
		zap.String("TraceFlags", sc.TraceFlags().String()),
		zap.String("TraceState", sc.TraceState().String()),
		zap.Bool("Remote", sc.IsRemote()),
	}
}
