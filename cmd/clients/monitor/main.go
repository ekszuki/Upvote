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
	if len(args) == 1 {
		logCtx.Panic("coin id is required")
	}

	coinID, err := strconv.ParseUint(args[1], 10, 32)
	if err != nil {
		logCtx.Panic("Invalid coin ID")
	}

	logCtx.Infof("starting new monitor to coin ID: %d", coinID)
	conn, err := grpc.Dial(
		"localhost:9090",
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
		logCtx.Errorf("could not open stream on gRPCServer: %v", err)
	}

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
				"Description: %s \nShort: %s \nVotes: %d",
				resp.GetDescription(),
				resp.GetShort(),
				resp.GetVotes(),
			)
		}
	}()

	<-done
	logCtx.Info("Stream finished")
}
