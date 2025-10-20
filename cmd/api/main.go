package main

import (
	"context"
	"database/sql"
	"flag"
	"log/slog"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Syha-01/national-inservice-training/internal/data"
	"github.com/Syha-01/national-inservice-training/internal/mailer"
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
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}

type application struct {
	config          configuration
	logger          *slog.Logger
	models          data.Models
	mailer          mailer.Mailer
	wg              sync.WaitGroup
	permissionModel data.PermissionModel
}

func main() {
	var settings configuration

	flag.IntVar(&settings.port, "port", 4000, "Server port")
	flag.StringVar(&settings.env, "env", "development",
		"Environment(development|staging|production)")
	// read in the dsn from the environment
	dsn := os.Getenv("TRAINING_DB_DSN")

	flag.StringVar(&settings.db.dsn, "db-dsn", dsn, "PostgreSQL DSN")

	// CORS configuration
	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		settings.cors.trustedOrigins = strings.Fields(val)
		return nil
	})

	// Rate limiter configuration
	flag.Float64Var(&settings.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&settings.limiter.burst, "limiter-burst", 5, "Rate limiter maximum burst")
	flag.BoolVar(&settings.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")

	// SMTP configuration
	flag.StringVar(&settings.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&settings.smtp.port, "smtp-port", 2525, "SMTP port")
	flag.StringVar(&settings.smtp.username, "smtp-username", "249f70e4cc4329", "SMTP username")
	flag.StringVar(&settings.smtp.password, "smtp-password", "2fa704cb93fea2", "SMTP password")
	flag.StringVar(&settings.smtp.sender, "smtp-sender", "National Inservice Training <no-reply@nits.com>", "SMTP sender")

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
		config:          settings,
		logger:          logger,
		models:          data.NewModels(db),
		mailer:          mailer.New(settings.smtp.host, settings.smtp.port, settings.smtp.username, settings.smtp.password, settings.smtp.sender),
		permissionModel: data.PermissionModel{DB: db},
	}

	err = appInstance.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
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
