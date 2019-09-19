package config

import "github.com/fwidjaya/demo-go-grpc-server/internal/models"

var LocalStorage = make([]*models.Transfer, 0)

const (
	HTTP_ADDR = ":8000"
	GRPC_SADDR = ":5000"
)
