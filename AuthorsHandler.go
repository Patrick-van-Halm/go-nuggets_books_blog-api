package main

import (
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/models"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func handleAuthorsGetRoutes(get *mux.Router) {
	get.HandleFunc("", authorsGetAllHandler)
	get.HandleFunc("/{id}", authorsGetHandler)
}

func handleAuthorsPostRoutes(post *mux.Router) {
	post.HandleFunc("", authorsPostHandler)
}

func authorsPostHandler(w http.ResponseWriter, r *http.Request) {
	handleCreate(&models.Author{}, w, r)
}

func authorsGetAllHandler(w http.ResponseWriter, _ *http.Request) {
	authors, err := models.GetAllAuthors(db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting all authors", zap.Error(err))
		return
	}

	handleJsonResponse(w, authors)
}

func authorsGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	author := models.Author{Id: id}
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