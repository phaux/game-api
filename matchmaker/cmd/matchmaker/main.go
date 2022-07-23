package main

import (
	"net/http"

	"github.com/phaux/game-api/matchmaker/gen/rpc/v1/v1connect"
	v1 "github.com/phaux/game-api/matchmaker/rpc/v1"
)

func main() {

	mux := http.NewServeMux()

	mux.Handle(v1connect.NewMatchmakerServiceHandler(v1.MatchmakerService{}))

	srv := &http.Server{
		Addr:    ":4452",
		Handler: mux,
	}

	srv.ListenAndServe()
}
