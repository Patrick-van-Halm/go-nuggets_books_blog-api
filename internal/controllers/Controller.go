package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/authenticator"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io"
	"net/http"
)
var logger *zap.Logger
func init() {
	logger = zap.L()
}

type Controller struct {
	Endpoints 	Endpoint
	Prefix    	string
	Database	*sql.DB
}

type Endpoint interface {
	setDatabase(db *sql.DB)
	create(w http.ResponseWriter, r *http.Request)
	readAll(w http.ResponseWriter, r *http.Request)
	read(w http.ResponseWriter, r *http.Request)
	update(w http.ResponseWriter, r *http.Request)
	delete(w http.ResponseWriter, r *http.Request)
}

func (c *Controller) MapEndpoints(r *mux.Router) {
	router := r.PathPrefix(c.Prefix).Subrouter()

	authRouter := router.NewRoute().Subrouter()
	authRouter.Use(authenticator.AuthorizationMiddleware)

	authRouter.HandleFunc("", c.Endpoints.create).Methods(http.MethodPost)
	router.HandleFunc("", c.Endpoints.readAll).Methods(http.MethodGet)
	router.HandleFunc("/{id}", c.Endpoints.read).Methods(http.MethodGet)
	authRouter.HandleFunc("/{id}", c.Endpoints.update).Methods(http.MethodPut, http.MethodPatch)
	authRouter.HandleFunc("/{id}", c.Endpoints.delete).Methods(http.MethodDelete)
	c.Endpoints.setDatabase(c.Database)
}

func parseJsonFromBody(body io.ReadCloser, data interface{}) error {
	decoder := json.NewDecoder(body)
	defer body.Close()
	return decoder.Decode(&data)
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