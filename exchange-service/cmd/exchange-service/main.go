package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/exchange-service/internal/channels/rest"
	"github.com/walbety/transaction-app/exchange-service/internal/config"
	"github.com/walbety/transaction-app/exchange-service/internal/service"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	log.Info("args:", os.Args)
	err := config.Initialize(os.Args)
	if err != nil {
		log.Fatal("error at initializing configs")
		os.Exit(2)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGINT)

	svc := service.New()

	log.Infof("%s starting at port: %s", config.Env.ServiceName, config.Env.RestPort)
	go func() {
		if err := rest.Start(svc); err != nil {
			log.WithError(err).Panic("error on http server")
		}
	}()

	<-stop
	rest.Stop(ctx)

	fmt.Print("aaaa\n\n")
}
