package main

import (
	"database/sql"
	"encoding/json"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/authenticator"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"io"
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
	handleBooksRoutes(r)
	handleReviewsRoutes(r)

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

func handleHttpError(w http.ResponseWriter, code int, message string, fields ...zap.Field) {
	logger.Error(message,
		fields...
	)
	http.Error(w, http.StatusText(code), code)
}

func writeJsonResponse(w http.ResponseWriter, json []byte) {
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(json); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst writing a response", zap.Error(err))
	}
}

func handleJsonResponse(w http.ResponseWriter, value interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst encoding json", zap.Error(err))
		return
	}

	writeJsonResponse(w, b)
}

func parseJsonFromBody(body io.ReadCloser, data interface{}) error {
	decoder := json.NewDecoder(body)
	defer body.Close()
	return decoder.Decode(&data)
}