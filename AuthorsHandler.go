package main

import (
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/classes"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func handleAuthorsRoutes(r *mux.Router){
	r.HandleFunc("/api/authors", authorsHandler).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/api/authors/{id}", authorGetHandler).Methods(http.MethodGet)
}

func authorsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		authorsGetAllHandler(w)
		break
	case http.MethodPost:
		authorsPostHandler(w, r.Body)
		break
	}
}

func authorsPostHandler(w http.ResponseWriter, body io.ReadCloser) {
	var data classes.Author
	if err := parseJsonFromBody(body, &data); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst parsing json", zap.Error(err))
		return
	}

	if err := data.New(db); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst creating a new author",
			zap.Error(err),
			zap.Any("author", &data),
		)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte("Created")); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst writing a response", zap.Error(err))
	}
}

func authorsGetAllHandler(w http.ResponseWriter) {
	authors, err := classes.GetAllAuthors(db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting all authors", zap.Error(err))
		return
	}

	handleJsonResponse(w, authors)
}

func authorGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	author := classes.Author{Id: id}
	err := author.Get(db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting specific author",
			zap.Error(err),
			zap.String("id", id),
		)
		return
	}

	handleJsonResponse(w, author)
}