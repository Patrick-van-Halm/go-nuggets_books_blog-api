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
	var err error

	// Setup logger
	if logger, err = zap.NewProduction(); err != nil {
		panic("failed to create logger, err: " + err.Error())
	}
	defer logger.Sync()

	// Setup environment variables using .env file
	if err = godotenv.Load(); err != nil {
		logger.Fatal("failed loading env file",
			zap.String("error", err.Error()),
		)
	}

	// Open connection to the database
	if db, err = sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING")); err != nil {
		logger.Fatal("failed connecting to database",
			zap.String("error", err.Error()),
		)
	}

	// Enable database pooling
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5*time.Minute)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/books", booksHandler).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/api/books/{id}", bookByIdHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/reviews", reviewsHandler).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/api/reviews/{id}", reviewByIdHandler).Methods(http.MethodGet)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(authenticator.AuthorizationMiddleware)

	if err := http.ListenAndServe(os.Getenv("HTTP_HOSTNAME"), r); err != nil {
		logger.Fatal("whilst listening an error occurred",
			zap.String("error", err.Error()),
		)
	}
}