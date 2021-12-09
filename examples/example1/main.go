package main

import (
	"errors"

	"github.com/xylonx/zapx"
	"go.uber.org/zap"
)

func main() {
	// using the embedded logger
	zapx.Info("Hello")
	zapx.Error("this is an error", zap.Error(errors.New("error!")))
	zapx.Warnf("%v", "warn")

	// using logger instance way
	logger, err := zapx.NewLogger(&zapx.Option{})
	if err != nil {
		panic(err)
	}

	logger.Info("hello")
	logger.Error("this is an error", zap.Error(errors.New("error!")))
	logger.Warnf("%v", "warn")

	// using overwritten embedded logger
	err = zapx.Use(&zapx.Option{})
	if err != nil {
		panic(err)
	}

	zapx.Info("Hello")
	zapx.Error("this is an error", zap.Error(errors.New("error!")))
	zapx.Warnf("%v", "warn")
}
