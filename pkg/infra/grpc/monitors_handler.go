package grpc

import (
	"context"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"klever.io/interview/pkg/application"
	protoMonitor "klever.io/interview/pkg/protos/monitors"
	"klever.io/interview/utils"
)

func (s *Server) CoinMonitorVotes(
	req *protoMonitor.CoinVoteMonitorRequest,
	stream protoMonitor.VoteMonitor_CoinMonitorVotesServer,
) error {
	logCtx := logrus.WithFields(
		logrus.Fields{"component": "gRPC Server", "method": "CoinMonitorVotes"},
	)
	logCtx.Info("starting stream request")
	refreshTime := getRefreshTime("VOTE_MONITOR_REFRESH_TIME_SEC")

	coinApp := application.NewCoinApplication(s.Repositories.Coins)
	var err error
	for {
		coin, err := coinApp.FindByID(context.Background(), uint(req.GetCoinID()))
		if err != nil {
			break
		}
		r := toCoinVoteMonitorResponse(coin)
		stream.Send(r)

		time.Sleep(time.Duration(refreshTime) * time.Second)
	}

	return err
}

func getRefreshTime(envName string) time.Duration {
	sRefreshTime := utils.GetEnv(envName, "5")
	refreshTime, err := strconv.Atoi(sRefreshTime)
	if err != nil {
		refreshTime = 5
	}

	return time.Duration(refreshTime)
}
