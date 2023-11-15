package main

import (
	"context"
	"fmt"
	"github.com/walbety/transaction-app/transaction-service/internal/integration"
	"github.com/walbety/transaction-app/transaction-service/internal/repository"

	log "github.com/sirupsen/logrus"

	"github.com/walbety/transaction-app/transaction-service/internal/channels/rest"
	"github.com/walbety/transaction-app/transaction-service/internal/config"
	"github.com/walbety/transaction-app/transaction-service/internal/service"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err := config.Initialize()
	if err != nil {
		log.Fatal("error at initializing configs")
		os.Exit(2)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGINT)

	exchange := integration.NewExchangeService()
	persist, err := repository.NewClient(ctx)
	if err != nil {
		log.Panic("ERRROORRRORORO")
	}

	svc := service.New(exchange, persist)

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
