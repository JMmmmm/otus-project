package sender

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Logger LoggerConf
	RMQ    RmqConfig
}

type LoggerConf struct {
	Level      string
	OutputPath string
}

type RmqConfig struct {
	URI          string
	Exchange     string
	ExchangeType string
	Reliable     bool
	Key          string
	Queue        string
}

func NewConfig(fpath string) (c Config, err error) {
	_, err = toml.DecodeFile(fpath, &c)
	return
}
