package main

import (
	"context"
	"flag"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/logger"
	"log"
	"os/signal"
	"syscall"

	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/app/sender"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/sender_config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()
	config, err := sender.NewConfig(configFile)
	if err != nil {
		log.Panicf("Can not read config: %s, %v", configFile, err)
	}
	logg, err := logger.NewAppLogger(config.Logger.Level, config.Logger.OutputPath)
	if err != nil {
		log.Panicf("Can not create logger: %v", err)
	}

	worker, err := sender.CreateWorker(ctx, config, logg)
	if err != nil {
		log.Panicf("Can not create worker: %v", err)
	}

	err = worker.Execute(ctx, 1)
	if err != nil {
		log.Panicf("Worker can not execute : %v", err)
	}
}
