package grpc

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/exchange-service/internal/channels/grpc/impl"
	"github.com/walbety/transaction-app/exchange-service/internal/config"
	"github.com/walbety/transaction-app/exchange-service/internal/service"
	"google.golang.org/grpc"
	"net"
)

type grpcServer struct {
	svc service.Service
	impl.UnimplementedExchangeServiceServer
}

func Listen(ctx context.Context, svc service.Service) *grpc.Server {
	port := config.Env.GrpcPort
	logc := log.WithContext(ctx).WithField("port", port)
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logc.WithError(err).Panic("error starting GRPC listener")
	}
	server := grpc.NewServer()
	impl.RegisterExchangeServiceServer(server, grpcServer{
		svc: svc,
	})

	go func() {
		err = server.Serve(listener)
		if err != nil {
			logc.WithError(err).Panic("error starting GRPC server")
		}
	}()
	logc.Info("Initialize GRPC server")

	return server

}
