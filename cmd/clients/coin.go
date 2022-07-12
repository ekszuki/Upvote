package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"klever.io/interview/pkg/protos/coins"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not create connection with grpc server %v", err)
	}
	grpcCli := coins.NewCoinServiceClient(conn)

	resp, err := grpcCli.CreateCoin(
		context.Background(),
		&coins.CreateCoinRequest{
			Description: "Coin 1",
			Short:       "C1",
		},
	)

	if err != nil {
		fmt.Printf("could not make request to server: %v", err)
	} else {

		fmt.Printf("Coin created: %v", resp)
	}
}
