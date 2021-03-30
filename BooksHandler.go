package main

import (
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/classes"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func handleBooksRoutes(r *mux.Router){
	r.HandleFunc("/api/books", booksHandler).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/api/books/{id}", bookGetHandler).Methods(http.MethodGet)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		booksHandlerGetAll(w)
		break
	case http.MethodPost:
		booksHandlerPost(w, r.Body)
		break
	}
}

func booksHandlerPost(w http.ResponseWriter, body io.ReadCloser) {
	var data classes.Book
	if err := parseJsonFromBody(body, &data); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst parsing json", zap.Error(err))
		return
	}

	if err := data.New(db); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst creating a new book",
			zap.Error(err),
			zap.Any("book", &data),
		)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte("Created")); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst writing a response", zap.Error(err))
	}
}

func booksHandlerGetAll(w http.ResponseWriter) {
	books, err := classes.GetAllBooks(db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting all books", zap.Error(err))
		return
	}

	handleJsonResponse(w, books)
}

func bookGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	book := classes.Book{Id: id}
	err := book.Get(db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting specific book",
			zap.Error(err),
			zap.String("id", id),
		)
		return
	}

	handleJsonResponse(w, book)
}