package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"
	"time"

	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/app/calendar"
	internalhttp "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/server/http"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/logger"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/calendar_config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config, err := calendar.NewConfig(configFile)
	if err != nil {
		log.Fatalf("Can not read config: %s, %v", configFile, err)
	}
	logg, err := logger.NewAppLogger(config.Logger.Level, config.Logger.OutputPath)
	if err != nil {
		log.Fatalf("Can not create logger: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	app, err := calendar.CreateApp(ctx, config, logg)
	if err != nil {
		logg.Error(fmt.Sprintf("can not create App: %v", err))
		return
	}
	server := internalhttp.NewServer(logg, app, net.JoinHostPort(config.Server.Host, config.Server.Port))

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		err := server.Stop(ctx)
		cancel()
		if err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		return
	}
}
