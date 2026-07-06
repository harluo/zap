package config

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newLogging, // 配置
	).Build().Apply()
}
