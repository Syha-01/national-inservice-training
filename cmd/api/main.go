package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

const version = "1.0.0"

type configuration struct {
	port int
	env  string
	cors struct {
		trustedOrigins []string
	}
}

type application struct {
	config configuration
	logger *slog.Logger
}

func main() {
	var settings configuration

	flag.IntVar(&settings.port, "port", 4000, "Server port")
	flag.StringVar(&settings.env, "env", "development",
		"Environment (development|staging|production)")
	
	// CORS configuration
	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		settings.cors.trustedOrigins = strings.Fields(val)
		return nil
	})
	
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	appInstance := &application{
		config: settings,
		logger: logger,
	}

	// Create server with proper configuration
	apiServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", settings.port),
		Handler:      appInstance.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server",
		"address", apiServer.Addr,
		"environment", settings.env,
		"cors", settings.cors.trustedOrigins,
	)

	err := apiServer.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}