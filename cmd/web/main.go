package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"
const cssVersion = "1"

// create a type config information of our applciation

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
	}
}

// config type for data items

type application struct {
	config config
	// loggers
	infoLog  *log.Logger
	errorLog *log.Logger
	// template cache
	templateCache map[string]*template.Template
	version       string
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
	app.infoLog.Println("Starting HTTP server in %s mode  on port %d", app.config.env, app.config.port)

	return srv.ListenAndServe()

}

func main() {
	// create a variable of type config
	// populate variable with information

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "port", "development", "Application environment {development|production}")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to api")
	flag.Parse()
	// get this stripe secret key and stripe publishable key
	//
	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	tc := make(map[string]*template.Template)

	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
	}
	// create the web server

}
