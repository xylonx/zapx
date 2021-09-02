package zapxdecoder

import (
	"context"
	"fmt"

	"github.com/SkyAPM/go2sky"
	"go.uber.org/zap"
)

var (
	SkywalkingDecoder skywalkingDecoder
)

type skywalkingDecoder struct{}

func (decoder skywalkingDecoder) DecodeCtx(ctx context.Context) []zap.Field {
	serviceName := go2sky.ServiceName(ctx)
	serviceInstanceName := go2sky.ServiceInstanceName(ctx)
	traceId := go2sky.TraceID(ctx)
	traceSegId := go2sky.TraceSegmentID(ctx)
	spanId := go2sky.SpanID(ctx)
	return []zap.Field{
		zap.String("SW_CTX", fmt.Sprintf("[%s,%s,%s,%s,%d]",
			serviceName, serviceInstanceName, traceId, traceSegId, spanId,
		)),
	}
}
