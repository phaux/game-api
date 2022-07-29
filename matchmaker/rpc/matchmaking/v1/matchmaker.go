package v1

import (
	"context"

	"github.com/bufbuild/connect-go"

	v1 "github.com/phaux/game-api/matchmaker/gen/rpc/matchmaking/v1"
	"github.com/phaux/game-api/matchmaker/gen/rpc/matchmaking/v1/v1connect"
)

type MatchmakerService struct{}

var _ v1connect.MatchmakerServiceHandler = MatchmakerService{}

func (MatchmakerService) FindMatch(
	ctx context.Context,
	req *connect.Request[v1.FindMatchRequest],
) (
	resp *connect.Response[v1.FindMatchResponse],
	err error,
) {
	return connect.NewResponse(&v1.FindMatchResponse{MatchId: "match-id"}), nil
}

func (MatchmakerService) ReportMatchResult(
	ctx context.Context,
	req *connect.Request[v1.ReportMatchResultRequest],
) (
	*connect.Response[v1.ReportMatchResultResponse],
	error,
) {
	return connect.NewResponse(&v1.ReportMatchResultResponse{}), nil
}
