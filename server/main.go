package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Kittonn/go-grpc/proto/currency"
	"google.golang.org/grpc"
)

const (
	port = ":9092"
)

type server struct {
	pb.UnimplementedCurrencyServer
}

func (s *server) GetRate(ctx context.Context, rateRequest *pb.RateRequest) (*pb.RateReply, error) {
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
