# zapx

this is a wrapper of [zap](https://github.com/uber-go/zap). The biggest difference of them is that the wrapper add `WithContext` method for logger.

**Attention: Sugger is not supportted now!**

## How to Use

Wrap the zap logger by `zapx.WrapZapLogger()`

the `zapx.WrapZapLogger()` function receives 2 args: logger which is `*zap.Logger`, and decoder which is an interface `CtxDecoder`.

the `CtxDecoder` interface contains just one method: `DecodeCtx(context.Context) []zap.Field`. It gets values from context and converts them into zap.Field. If passing nil, it will ignore context.

If you want to get value from context, you can define your own decoder and pass it as the second parameter.

examples:

- decoder is nil:

```go
func main() {
	l := zap.NewExample()
	logger := zapx.WrapZapLogger(l, nil)
	ctx := context.WithValue(context.Background(), "A", "ababab")
	logger.WithContext(ctx).Info("xxxx")
}
```

- self-define decoder:

```go
type traceDecoder struct{}

func (decoder traceDecoder) DecodeCtx(ctx context.Context) []zap.Field {
	fields := make([]zap.Field, 0)
	if traceId, ok := ctx.Value("TraceId").(string); ok {
		fields = append(fields, zap.String("TraceId", traceId))
	}
	return fields
}

func main() {
	l := zap.NewExample()
	logger := zapx.WrapZapLogger(l, traceDecoder{})
	ctx := context.WithValue(context.Background(), "TraceId", "ababab")
	logger.WithContext(ctx).Info("xxxx")
}
```

or using the pre-wrapped function

1. init the zapx

```go
func init() {
	zapx.WrapZapLogger(zap.NewExample(), nil)

	zapx.Info("Hello", zap.String("env", "test"))
}
```
