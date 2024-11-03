package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Kittonn/go-grpc/proto/currency"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
		Destination: pb.Currencies_CZK,
	})

	if err != nil {
		log.Fatalf("Can't get rate: %v", err)
	}

	log.Printf("Rate: %v", r.GetRate())
}
