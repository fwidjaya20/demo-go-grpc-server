package containers

import "github.com/fwidjaya20/demo-go-grpc-server/internal"

type ServiceContainer struct {
	TransferService internal.ServiceInterface
}

func NewServiceContainer() ServiceContainer {
	return ServiceContainer{
		TransferService: internal.NewTransferService(),
	}
}