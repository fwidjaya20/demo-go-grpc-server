package internal

import (
	"github.com/fwidjaya20/demo-go-grpc-server/config"
	"github.com/fwidjaya20/demo-go-grpc-server/internal/models"
)

type ServiceInterface interface {
	GetHistory() ([]*models.Transfer, error)
	Transfer(payload *models.Transfer) (*models.Transfer, error)
}

type service struct {}

func NewTransferService() ServiceInterface {
	return &service{}
}

func (t *service) GetHistory() ([]*models.Transfer, error) {
	return config.LocalStorage, nil
}

func (t *service) Transfer(payload *models.Transfer) (*models.Transfer, error) {
	config.LocalStorage = append(config.LocalStorage, payload)

	return payload, nil
}
