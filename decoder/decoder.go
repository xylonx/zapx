package decoder

import (
	"context"

	"go.uber.org/zap"
)

type CtxDecoder interface {
	DecodeCtx(context.Context) []zap.Field
}

type NoopCtxDecoder struct{}

func (decoder *NoopCtxDecoder) DecodeCtx(ctx context.Context) []zap.Field {
	return nil
}
