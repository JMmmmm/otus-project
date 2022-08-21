package scheduler

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Logger               LoggerConf
	PSQL                 PSQLConfig
	RMQ                  RmqConfig
	DB                   DB
	SchedulerInterval    int
	NotificationInterval int
}

type LoggerConf struct {
	Level      string
	OutputPath string
}

type PSQLConfig struct {
	DSN string
}

type RmqConfig struct {
	URI          string
	Exchange     string
	ExchangeType string
	Reliable     bool
	RoutingKey   string
	Queue        string
}

type DB struct {
	DBType string
}

func NewConfig(fpath string) (c Config, err error) {
	_, err = toml.DecodeFile(fpath, &c)
	return
}
