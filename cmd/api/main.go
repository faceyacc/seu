package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

const version = "1.0.0"

// config holds all configurtion settings
// for the appication.
type config struct {
	port int
	env  string
}

// application holds dependencies for HTTP handlers
// helpers, and middleware.
type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "SeuAPI server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (dev|stage|prod)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
