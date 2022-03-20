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

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/app"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/server/http"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config, err := app.NewConfig(configFile)
	if err != nil {
		log.Fatalf("Can not read config: %s, %v", configFile, err)
	}
	logg, err := logger.NewAppLogger(config.Logger.Level, config.Logger.OutputPath)
	if err != nil {
		log.Fatalf("Can not create logger: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	calendar, err := app.CreateApp(ctx, config, logg)
	if err != nil {
		logg.Error(fmt.Sprintf("can not create App: %v", err))
		return
	}
	server := internalhttp.NewServer(logg, calendar, net.JoinHostPort(config.Server.Host, config.Server.Port))

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
