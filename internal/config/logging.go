package config

import (
	"github.com/harluo/config"
)

type Logging struct {
	// 日志级别
	Level string `default:"debug" json:"level,omitempty" validate:"oneof=debug info warn error fatal"`
	// 日志调用方法过滤层级
	Skip int `default:"2" json:"skip,omitempty"`
	// 调用堆栈层级
	Stacktrace *int `json:"stacktrace,omitempty"`
}

func newLogging(getter config.Getter) (logging *Logging, err error) {
	logging = new(Logging)
	err = getter.Get(&struct {
		Logging *Logging `json:"logging,omitempty" validate:"required"`
	}{
		Logging: logging,
	})

	return
}
