package config

import (
	"github.com/koding/multiconfig"
)

type (
	ServerConfig struct {
		Name    string    `default:"gopher"`
		Host    string    `default:":6060"`
		Database Database
	}

	Database struct {
		Driver    string
		Source    string
	}
)

func GetConfig() *ServerConfig  {
	//读取配置
	m := multiconfig.NewWithPath("config/config.toml")
	c := &ServerConfig{}
	m.MustLoad(c)

	return c
}
