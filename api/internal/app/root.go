package app

import (
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
)

type Root struct {
	Server      ServerConfig  `mapstructure:"server"`
	Log         log.Config    `mapstructure:"log"`
	MiddleWare  mid.LogConfig `mapstructure:"middleware"`
	Credentials string        `mapstructure:"credentials"`
}

type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int64  `mapstructure:"port"`
}
