package main

import (
	"database/sql"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/authenticator"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

var db *sql.DB
var logger *zap.Logger

func init() {
	// Setup logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic("failed to create logger, err: " + err.Error())
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	// Setup environment variables using .env file
	if err = godotenv.Load(); err != nil {
		logger.Fatal("failed loading env file",
			zap.Error(err),
		)
	}

	// Open connection to the database
	connString := getEnv("DB_CONNECTION_STRING")
	if db, err = sql.Open("postgres", connString); err != nil {
		logger.Fatal("failed connecting to database",
			zap.Error(err),
			zap.String("connection_string", connString),
		)
	}

	// Enable database pooling
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5*time.Minute)
}

func main() {
	logger = zap.L()

	r := mux.NewRouter()
	r.HandleFunc("/api/books", booksHandler).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/api/books/{id}", bookByIdHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/reviews", reviewsHandler).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/api/reviews/{id}", reviewByIdHandler).Methods(http.MethodGet)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(authenticator.AuthorizationMiddleware)

	// Start webserver
	hostname := getEnv("HTTP_HOSTNAME")
	err := http.ListenAndServe(hostname, r)
	if err != nil {
		logger.Fatal("failed to start listening",
			zap.Error(err),
			zap.String("hostname", hostname),
		)
	}
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		logger.Fatal("environment variable not set",
			zap.String("key", key),
		)
	}
	return value
}