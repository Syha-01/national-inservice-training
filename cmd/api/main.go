package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const appVersion = "1.0.0"

type configuration struct {
	port int
	env  string
	db   struct {
		dsn string
	}
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
		"Environment(development|staging|production)")
	// read in the dsn
	flag.StringVar(&settings.db.dsn, "db-dsn", "postgres://nits:bananaforscale@localhost/nits?sslmode=disable", "PostgreSQL DSN")

	// CORS configuration
	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		settings.cors.trustedOrigins = strings.Fields(val)
		return nil
	})

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// the call to openDB() sets up our connection pool
	db, err := openDB(settings)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	// release the database resources before exiting
	defer db.Close()
	logger.Info("database connection pool established")

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

	logger.Info("starting server", "address", apiServer.Addr,
		"environment", settings.env)
	err = apiServer.ListenAndServe() // remove the :
	logger.Error(err.Error())
	os.Exit(1)
} // end of main()

func openDB(settings configuration) (*sql.DB, error) {
	// open a connection pool
	db, err := sql.Open("postgres", settings.db.dsn)
	if err != nil {
		return nil, err
	}

	// set a context to ensure DB operations don't take too long
	ctx, cancel := context.WithTimeout(context.Background(),
		5*time.Second)
	defer cancel()
	// let's test if the connection pool was created
	// we trying pinging it with a 5-second timeout
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}
	// return the connection pool (sql.DB)
	return db, nil
}
