package http

import (
	"fmt"
	"github.com/fwidjaya/demo-go-grpc-server/config"
	"log"
	netHttp "net/http"

	"github.com/fwidjaya/demo-go-grpc-server/cmd/containers"
	"github.com/fwidjaya/demo-go-grpc-server/internal/transports/http"
	"github.com/go-chi/chi"
)

// HTTPServer .
func HTTPServer(containers containers.ServiceContainer) {
	var router *chi.Mux

	router = chi.NewRouter()

	router.Post("/transfer", http.Transfer(containers.TransferService))
	router.Get("/history", http.GetHistory(containers.TransferService))

	log.Println(fmt.Sprintf("HTTP Transport running on %s", config.HTTP_ADDR))
	_ = netHttp.ListenAndServe(config.HTTP_ADDR, router)
}