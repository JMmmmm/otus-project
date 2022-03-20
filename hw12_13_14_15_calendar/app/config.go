package app

import "github.com/BurntSushi/toml"

type Config struct {
	Logger LoggerConf
	PSQL   PSQLConfig
	Server ServerConfig
	DB     DB
}

type LoggerConf struct {
	Level      string
	OutputPath string
}

type PSQLConfig struct {
	DSN string
}

type ServerConfig struct {
	Host string
	Port string
}

type DB struct {
	DBType string
}

func NewConfig(fpath string) (c Config, err error) {
	_, err = toml.DecodeFile(fpath, &c)
	return
}
