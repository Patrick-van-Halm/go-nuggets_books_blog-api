package main

import (
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/models"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func handleGenresGetRoutes(get *mux.Router) {
	get.HandleFunc("", genresGetAllHandler)
	get.HandleFunc("/{id}", genreGetHandler)
}

func handleGenresPostRoutes(post *mux.Router) {
	post.HandleFunc("", genresPostHandler)
}

func genresPostHandler(w http.ResponseWriter, r *http.Request) {
	handleCreate(&models.Genre{}, w, r)
}

func genresGetAllHandler(w http.ResponseWriter, _ *http.Request) {
	genres, err := models.GetAllGenres(db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting all genres", zap.Error(err))
		return
	}

	handleJsonResponse(w, genres)
}

func genreGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	genre := models.Genre{Id: id}
	err := genre.Get(db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting specific genre",
			zap.Error(err),
			zap.String("id", id),
		)
		return
	}

	handleJsonResponse(w, genre)
}