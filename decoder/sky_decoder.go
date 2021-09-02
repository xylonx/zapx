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
	return []zap.Field{
		zap.String("SW_CTX", fmt.Sprintf("[%s,%s,%s,%s,%d]",
			go2sky.ServiceName(ctx),
			go2sky.ServiceInstanceName(ctx),
			go2sky.TraceID(ctx),
			go2sky.TraceSegmentID(ctx),
			go2sky.SpanID(ctx),
		)),
	}
}
