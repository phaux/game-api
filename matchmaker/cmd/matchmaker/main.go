package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/phaux/game-api/matchmaker/gen/rpc/matchmaking/v1/v1connect"
	v1 "github.com/phaux/game-api/matchmaker/rpc/matchmaking/v1"
)

func main() {

	setUpLogLevel()

	if !tzDataOK() {
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.Handle(v1connect.NewMatchmakerServiceHandler(v1.MatchmakerService{}))

	runServer(mux)
}

func setUpLogLevel() {
	rawLvl := os.Getenv("LOG_LEVEL")

	log.Logger = log.Logger.Level(zerolog.InfoLevel)

	lvl, err := zerolog.ParseLevel(rawLvl)

	if err != nil {
		log.Warn().
			Str("env.LOG_LEVEL", rawLvl).
			Err(err).
			Msg("cannot parse specified log level")
		return
	}

	if lvl == zerolog.NoLevel {
		lvl = zerolog.InfoLevel
	}

	log.Info().
		Str("log.level", lvl.String()).
		Msg("setting log level")

	log.Logger = log.Logger.Level(lvl)
}

func tzDataOK() bool {
	name := "Europe/Warsaw"

	_, err := time.LoadLocation(name)
	if err != nil {
		log.Error().
			Str("test.location", name).
			Err(err).
			Msg("cannot load the test location from the available timezone information")
	}

	return err == nil
}

func runServer(h http.Handler) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4447"
	}

	addr := fmt.Sprintf(":%s", port)

	srv := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	log.Info().
		Str("addr", srv.Addr).
		Msg("starting server")

	srv.ListenAndServe()
}
