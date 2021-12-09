# zapx

this is a wrapper of [zap](https://github.com/uber-go/zap). The biggest difference of them is that the wrapper add `WithContext` method for logger.

## How to Use

there are 3 ways to use zapx:

- packaged-embedded way

nothing to init! just using the method `zapx.Info()`, `zapx.Warnf()` and so on.

```golang
func main() {
	zapx.Info("Hello")
	zapx.Error("this is an error", zap.Error(errors.New("error!")))
	zapx.Warnf("%v", "warn")
}
```

- global logger instance

```golang
func main() {
	logger, err := zapx.NewLogger(&zapx.Option{})
	if err != nil {
		panic(err)
	}

	logger.Info("hello")
	logger.Error("this is an error", zap.Error(errors.New("error!")))
	logger.Warnf("%v", "warn")
}
```

- overwrite packaged-embedded variable

```golang
func main() {
	err = zapx.Use(&zapx.Option{})
	if err != nil {
		panic(err)
	}

	zapx.Info("Hello")
	zapx.Error("this is an error", zap.Error(errors.New("error!")))
	zapx.Warnf("%v", "warn")
}
```

## Feature

context decoder

the CtxDecoder is just an interface:

```golang
type CtxDecoder interface {
	DecodeCtx(context.Context) []zap.Field
}
```

it 'decode' the context and get values from it. Then it converts them into []zap.Field to log. 

It is designed to integrated with `Tracing`: get tracing related info from context, like traceID.

Now, zapx implements some context decoder:

- [x] [open-telemetry](https://github.com/open-telemetry/opentelemetry-go)
- [x] [sky-walking](https://github.com/SkyAPM/go2sky)
- [ ] [elastic-apm](https://github.com/elastic/apm-agent-go)