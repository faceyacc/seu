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

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: mux,
	}

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version %s\n", version)
}