package grpc

import (
	"context"
	"github.com/fwidjaya20/demo-go-grpc-server/internal"
	"github.com/fwidjaya20/demo-go-grpc-server/internal/models"
	"log"

	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/fwidjaya20/demo-go-grpc-server/pkg/atm"

)

type GrpcService struct {
	TransferService internal.ServiceInterface
}

func (s GrpcService) Send(context context.Context, payload *pb.Transfer) (*pb.TransferResponse, error) {
	log.Println("[SEND]", payload.String())

	m := &models.Transfer{
		payload.Sender,
		payload.Recipient,
		payload.Amount,
	}

	result, err := s.TransferService.Transfer(m)

	if nil != err {
		return nil, err
	}

	return &pb.TransferResponse{
		Sender:               result.Sender,
		Recipient:            result.Recipient,
		Amount:               result.Amount,
	}, nil
}

func (s GrpcService) Receipt(context context.Context, void *empty.Empty) (*pb.TransferHistory, error) {
	log.Println("[Receipt]")

	results, err := s.TransferService.GetHistory()

	if nil != err {
		return nil, err
	}

	var response = &pb.TransferHistory{}

	for _, v := range results {
		response.List = append(response.List, &pb.Transfer{
			Sender:               v.Sender,
			Recipient:            v.Recipient,
			Amount:               v.Amount,
		})
	}

	return response, nil
}
