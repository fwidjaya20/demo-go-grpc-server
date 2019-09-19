package containers

import "github.com/fwidjaya/demo-go-grpc-server/internal"

type ServiceContainer struct {
	TransferService internal.ServiceInterface
}

func NewServiceContainer() ServiceContainer {
	return ServiceContainer{
		TransferService: internal.NewTransferService(),
	}
}