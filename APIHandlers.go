package main

import (
	"encoding/json"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/classes/Book"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/classes/Review"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			booksHandlerGet(w)
			break
		case http.MethodPost:
			booksHandlerPost(w, r.Body)
			break
	}
}

func booksHandlerPost(w http.ResponseWriter, body io.ReadCloser) {
	var data book.Book
	if err := parseJsonFromBody(body, &data); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst parsing json", zap.Error(err))
		return
	}

	if err := book.New(db, data); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst creating a new book", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte("Created")); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst writing a response", zap.Error(err))
	}
}

func booksHandlerGet(w http.ResponseWriter) {
	handleJsonResponse(w, book.GetAll(db))
}

func bookByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	handleJsonResponse(w, book.GetWithHash(db, vars["id"]))
}

func reviewsHandler(w http.ResponseWriter, _ *http.Request) {
	handleJsonResponse(w, review.GetAll(db))
}

func reviewByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	handleJsonResponse(w, review.GetWithHash(db, vars["id"]))
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