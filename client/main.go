package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Kittonn/go-grpc/proto/currency"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	creds := insecure.NewCredentials()

	conn, err := grpc.NewClient("localhost:9092", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	defer conn.Close()

	c := pb.NewCurrencyClient(conn)

	r, err := c.GetRate(ctx, &pb.RateRequest{
		Base: pb.Currencies_AUD,
		// Base: pb.Currencies(pb.Currencies_value["EUR"]),
		Destination: pb.Currencies_AUD,
	})

	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Printf("gRPC error code: %v, error message: %v", st.Code(), st.Message())
		} else {
			log.Fatalf("Unexpected error: %v", err)
		}

		return
	}

	log.Printf("Rate: %v", r.GetRate())
}
