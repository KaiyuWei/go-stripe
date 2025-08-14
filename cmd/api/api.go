package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
	}
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf(
		"Starting back end server in %s mode on port %d",
		app.config.env,
		app.config.port,
	)

	return srv.ListenAndServe()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	var cfg config

	flag.StringVar(
		&cfg.env,
		"env",
		"development",
		"Application environment {development|production|maintenance}",
	)

	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")
	cfg.db.dsn = os.Getenv("DSN")
	cfg.port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatal("Invalid API port:", err)
	}

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := sql.Open("mysql", cfg.db.dsn)
	if err != nil {
		errorLog.Fatal("Error connecting to the database:", err)
	}
	defer conn.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}
