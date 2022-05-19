package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"org.Magassians/routes"
	"org.Magassians/util"
)

var (
	PORT string
	wait time.Duration
)

func main() {
	PORT = os.Getenv("SERVER_PORT")

	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	server := &http.Server{
		Addr:         "localhost:" + PORT,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      util.HelmetHandler().Secure(routes.RouterIndex()),
	}

	go func() {
		util.Log.Info().Msg("Starting server on : " + PORT)
		if err := server.ListenAndServe(); err != nil {
			util.Log.Fatal().Err(err).Msg("Server Start failed")
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	server.Shutdown(ctx)
	util.Log.Info().Msg("shutting down")
	os.Exit(0)
}
