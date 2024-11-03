package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Kittonn/go-grpc/proto/currency"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":9092"
)

type server struct {
	pb.UnimplementedCurrencyServer
}

func (s *server) GetRate(ctx context.Context, rateRequest *pb.RateRequest) (*pb.RateReply, error) {
	log.Printf("Received request to convert from %s to %s", rateRequest.Base, rateRequest.Destination)

	if rateRequest.Base == rateRequest.Destination {
		return nil, status.Errorf(codes.InvalidArgument, "Base and destination are same")
	}

	return &pb.RateReply{
		Rate: 0.5,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCurrencyServer(grpcServer, &server{})

	log.Printf("gRPC server listening on %s", port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
