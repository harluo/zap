package core

import (
	"github.com/goexl/log"
	"github.com/harluo/zap/internal/config"
	"github.com/harluo/zap/internal/internal"
)

type Logger = log.Logger

func newLogger(config *config.Logging, factory *internal.Factory) (logger log.Logger, err error) {
	builder := log.New().Level(log.ParseLevel(config.Level))
	if nil != config.Stacktrace {
		builder.Stacktrace(*config.Stacktrace)
	}
	logger, err = builder.Factory(factory).Build()

	return
}
