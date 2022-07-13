package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"klever.io/interview/pkg/protos/monitors"
)

func main() {
	logCtx := logrus.WithFields(
		logrus.Fields{"component": "gRPC Monitor Client", "method": "main"},
	)

	args := os.Args
	if len(args) != 3 {
		fmt.Println("Invalid parameters, please use: ")
		fmt.Println(" main [server_url] [coinID]")
		fmt.Printf("\n Example: main localhost:9009 1")
		os.Exit(-1)
	}

	serverURL := args[1]
	if serverURL == "" {
		logCtx.Panic("Server URL is required")
	}

	coinID, err := strconv.ParseUint(args[2], 10, 32)
	if err != nil {
		logCtx.Panic("Invalid coin ID")
	}

	logCtx.Infof("connecting monitor on server %s", serverURL)
	conn, err := grpc.Dial(
		serverURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logCtx.Fatalf("could not connect on the gRPC Server")
	}

	grpcCli := monitors.NewVoteMonitorClient(conn)

	stream, err := grpcCli.CoinMonitorVotes(
		context.Background(),
		&monitors.CoinVoteMonitorRequest{
			CoinID: 2,
		},
	)
	if err != nil {
		logCtx.Fatalf("could not open stream on gRPCServer: %v", err)
	}

	logCtx.Infof("Monitoring coin id: %d", coinID)
	done := make(chan bool)
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				logCtx.Fatalf("could not receive response %v", err)
			}

			fmt.Printf(
				"Description: %s \nShort: %s \nVotes: %d\n",
				resp.GetDescription(),
				resp.GetShort(),
				resp.GetVotes(),
			)
		}
	}()

	<-done
	logCtx.Info("Stream finished")
}
