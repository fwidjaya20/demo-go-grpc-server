package grpc

import (
	"fmt"
	"github.com/fwidjaya20/demo-go-grpc-server/cmd/containers"
	"github.com/fwidjaya20/demo-go-grpc-server/config"
	grpcTransport "github.com/fwidjaya20/demo-go-grpc-server/internal/transports/grpc"
	pb "github.com/fwidjaya20/demo-go-grpc-server/pkg/atm"
	"google.golang.org/grpc"
	"log"
	"net"
)

func GRPCServer(containers containers.ServiceContainer) {
	grpcServer := grpc.NewServer()

	grpcService := &grpcTransport.GrpcService{
		TransferService: containers.TransferService,
	}

	pb.RegisterTransfersServer(grpcServer, grpcService)

	log.Println(fmt.Sprintf("GRPC Server runnging on %s", config.GRPC_SADDR))

	l, err := net.Listen("tcp", config.GRPC_SADDR)

	if nil != err {
		log.Fatalf("could not listen to %s: %v", config.GRPC_SADDR, err)
	}

	log.Fatal(grpcServer.Serve(l))
}