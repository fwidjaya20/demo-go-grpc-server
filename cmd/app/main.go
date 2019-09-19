package main

import (
	"fmt"
	"github.com/fwidjaya/demo-go-grpc-server/cmd/containers"
	"github.com/fwidjaya/demo-go-grpc-server/cmd/grpc"
	"github.com/fwidjaya/demo-go-grpc-server/cmd/http"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	svcContainer := containers.NewServiceContainer()

	runtime.GOMAXPROCS(2)
	go http.HTTPServer(svcContainer)
	go grpc.GRPCServer(svcContainer)

	var input string
	fmt.Scanln(&input)
}
