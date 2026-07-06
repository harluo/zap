package internal

import (
	"github.com/goexl/log"
	"github.com/harluo/zap/internal/internal/internal"
)

type Factory struct{}

func newFactory() *Factory {
	return new(Factory)
}

func (f *Factory) New() (log.Executor, error) {
	return internal.New()
}
