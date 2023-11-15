package grpc

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ConnectGprcInsecure(address string) *grpc.ClientConn {
	return ConnectGrpc(address, grpc.WithInsecure())
}

func ConnectGrpc(address string, options ...grpc.DialOption) *grpc.ClientConn {
	log.Infof("GRPC -> Connecting to %s", address)
	client, err := grpc.Dial(address, options...)
	if err != nil {
		log.Panic("Failed when connecting into ExchangeService")
	}
	return client
}

