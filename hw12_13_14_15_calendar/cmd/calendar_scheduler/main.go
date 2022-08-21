package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/logger"

	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/app/scheduler"
	"github.com/procyon-projects/chrono"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/scheduler_config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()
	config, err := scheduler.NewConfig(configFile)
	if err != nil {
		log.Fatalf("Can not read config: %s, %v", configFile, err)
	}
	logg, err := logger.NewAppLogger(config.Logger.Level, config.Logger.OutputPath)
	if err != nil {
		log.Fatalf("Can not create logger: %v", err)
	}

	worker, err := scheduler.CreateWorker(ctx, config, logg)
	if err != nil {
		log.Fatalf("Can not create worker: %v", err)
	}
	timeTo := time.Now()
	timeFrom := timeTo.Add(-(time.Duration(config.NotificationInterval) * time.Hour))

	taskScheduler := chrono.NewDefaultTaskScheduler()
	_, err = taskScheduler.ScheduleAtFixedRate(func(ctx context.Context) {
		worker.Execute(config.RMQ.Uri, config.RMQ.ExchangeType, config.RMQ.Queue, config.RMQ.Exchange, config.RMQ.RoutingKey, config.RMQ.Reliable, timeFrom, timeTo)
	}, time.Duration(config.SchedulerInterval)*time.Second)

	if err != nil {
		logg.Error(fmt.Sprintf("exchange Publish: %v", err))
	}

	<-ctx.Done()
	logg.Info("Scheduler stopped")
}
